package task

import (
	"context"
	"database/sql"
	"os"
	"testing"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var testDB *sql.DB

func TestMain(m *testing.M) {
	godotenv.Load("../../.env")

	dbUrl := os.Getenv("DB_URL")

	db, err := sql.Open("postgres", dbUrl)
	if err != nil {
		panic("cannot connect: " + err.Error())
	}

	err = db.Ping()
	if err != nil {
		panic("cannot connect: " + err.Error())
	}

	testDB = db
	m.Run()
	db.Close()
}

func TestCreateTaskRepo(t *testing.T) {
	repo := NewTaskRepository(testDB)

	var userID uuid.UUID
	testDB.QueryRow(
		"INSERT INTO users (name, email, password) VALUES ($1, $2, $3) RETURNING id",
		"Test User", "EvilRick@test.com", "Rick",
	).Scan(&userID)

	task := Task{
		Title:       "Test_title",
		Description: "Some description",
		AssignedBy:  userID,
		AssignedTo:  userID,
		Estimate:    1,
		Status:      "to_do",
	}

	ctx := context.Background()
	created, err := repo.CreateTask(ctx, task)

	if err != nil {
		t.Fatalf("didn't expect an error: %v", err)
	}

	if created.ID == (uuid.UUID{}) {
		t.Errorf("expected a completed ID")
	}

	t.Cleanup(func() {
		testDB.Exec("DELETE FROM tasks WHERE id = $1", created.ID)
		testDB.Exec("DELETE FROM users WHERE id = $1", userID)
	})
}

func TestGetByIDRepo(t *testing.T) {
	repo := NewTaskRepository(testDB)

	var userID uuid.UUID
	testDB.QueryRow(
		"INSERT INTO users (name, email, password) VALUES ($1, $2, $3) RETURNING id",
		"Rick", "EvilMorty@test.com", "Rick",
	).Scan(&userID)

	var task Task
	testDB.QueryRow(
		"INSERT INTO tasks (title, description, assigned_by, assigned_to, estimate) VALUES ($1, $2, $3, $4, $5) RETURNING created_at, id, status, updated_at, completed_at",
		"Assemble an intergalactic antimatter condenser", "We need to fly to Blips and Cheets, get three Class IX crystals, and not lose a single one. Don't screw up, Morty.", userID, userID, 3,
	).Scan(&task.CreatedAt, &task.ID, &task.Status, &task.UpdatedAt, &task.CompletedAt)

	ctx := context.Background()
	search, err := repo.GetByID(ctx, task.ID)

	if err != nil {
		t.Fatalf("didn't expect an error: %v", err)
	}

	if search.ID == (uuid.UUID{}) {
		t.Errorf("expected a completed ID")
	}

	t.Cleanup(func() {
		testDB.Exec("DELETE FROM tasks WHERE id = $1", search.ID)
		testDB.Exec("DELETE FROM users WHERE id = $1", userID)
	})

}

func TestUpdateTaskRepo(t *testing.T) {
	repo := NewTaskRepository(testDB)

	var userID uuid.UUID
	err := testDB.QueryRow(
		"INSERT INTO users (name, email, password) VALUES ($1, $2, $3) RETURNING id",
		"Rick", "MortySanchez@test.com", "Rick",
	).Scan(&userID)
	if err != nil {
		t.Fatalf("не смогли создать пользователя: %v", err)
	}

	var task Task
	testDB.QueryRow(
		"INSERT INTO tasks (title, description, assigned_by, assigned_to, estimate) VALUES ($1, $2, $3, $4, $5) RETURNING created_at, id, status, updated_at, completed_at",
		"Assemble an intergalactic antimatter condenser", "We need to fly to Blips and Cheets, get three Class IX crystals, and not lose a single one. Don't screw up, Morty.", userID, userID, 3,
	).Scan(&task.CreatedAt, &task.ID, &task.Status, &task.UpdatedAt, &task.CompletedAt)

	task.Title = "Fix megaseeds admin interface"
	task.Description = "The 'Destroy All' button is back on the home page. Remove it from the settings and add a confirmation. Well... no, let there be a confirmation."
	task.AssignedBy = userID
	task.AssignedTo = userID

	ctx := context.Background()
	update, err := repo.UpdateTask(ctx, task)

	if err != nil {
		t.Fatalf("didn't expect an error: %v", err)
	}

	if update.ID == (uuid.UUID{}) {
		t.Errorf("expected a completed ID")
	}

	t.Cleanup(func() {
		testDB.Exec("DELETE FROM tasks WHERE id = $1", update.ID)
		testDB.Exec("DELETE FROM users WHERE id = $1", userID)
	})
}

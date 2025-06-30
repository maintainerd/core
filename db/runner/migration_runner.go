package runner

import (
	"log"
	"os/exec"
)

func RunDefaultMigrations(connectionString string) {
	log.Println("🏃 Running migrations...")

	// Migration files
	directories := []string{
		"./db/migration/v1/core",
		"./db/migration/v1/auth",
	}

	// Run migrations
	for _, dir := range directories {
		runMigrationDir(dir, connectionString)
	}

	log.Println("✅ Migrations completed.")
}

func runMigrationDir(dir string, connectionString string) {
	cmd := exec.Command("goose", "-dir", dir, "postgres", connectionString, "up")
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("❌ Migration failed in %s: %v\nOutput: %s", dir, err, string(output))
	}
	log.Printf("✅ Migrations applied from %s", dir)
}

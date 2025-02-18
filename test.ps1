# test.ps1

Write-Host "=== Starting Docker containers ==="
docker compose up -d

Write-Host "=== Waiting for containers to become healthy (if they have a healthcheck) ==="

# List the container names from docker-compose.yml
$containers = @("mongo_test", "mysql_test", "postgres_test")

# Poll each container's Health.Status until all are "healthy" or we time out.
# If a container doesn't define a healthcheck, we treat it as "healthy" automatically.
$maxAttempts = 30  # 30 attempts, 5 seconds each = up to ~2.5 minutes
$attempt = 1

while ($true) {
    $allHealthy = $true

    foreach ($container in $containers) {
        # Retrieve the container's .State.Health.Status (if any)
        $health = docker inspect --format='{{.State.Health.Status}}' $container 2>$null

        # If .State.Health doesn't exist (no healthcheck), $health will be empty or null
        if (!$health) {
            Write-Host "Container $container has no healthcheck. Considering it healthy."
            continue
        }

        # If the container has a healthcheck but isn't yet 'healthy', wait
        if ($health -ne "healthy") {
            Write-Host "Container $container is '$health'. Waiting..."
            $allHealthy = $false
            break
        }
    }

    if ($allHealthy) {
        Write-Host "All containers that have healthchecks are healthy (others have no healthcheck)."
        break
    }

    if ($attempt -ge $maxAttempts) {
        Write-Host "Timed out waiting for containers to become healthy."
        exit 1
    }

    $attempt++
    Start-Sleep -Seconds 5
}

# Write-Host "=== Initializing MongoDB Test Collection ==="
# & mongosh --host localhost --port 27017 --eval "use testdb; db.createCollection('testCollection')"

Write-Host "=== Setting environment variables for tests ==="
# Use 127.0.0.1 if you find 'localhost' is problematic on Windows,
# plus a 20s connect timeout for Mongo.
$env:MONGO_URI     = "mongodb://localhost:27017"
$env:MYSQL_DSN     = "root:root@tcp(localhost:3306)/testdb?charset=utf8mb4&parseTime=True&loc=Local"
$env:POSTGRES_DSN  = "host=localhost port=5432 user=test password=test dbname=test sslmode=disable"

Write-Host "=== Running go test -v ./... ==="
go test -v ./test/./...

Write-Host "=== Done ==="

# Optional: If you get an execution policy error, you may need to allow local script execution:
# Set-ExecutionPolicy RemoteSigned -Scope CurrentUser
# Test script for the Go Fiber API

Write-Host "Testing Go Fiber API..." -ForegroundColor Green

# Test root endpoint
Write-Host "`n1. Testing root endpoint..." -ForegroundColor Yellow
try {
    $response = Invoke-RestMethod -Uri "http://localhost:8080/" -Method GET
    Write-Host "✓ Root endpoint working" -ForegroundColor Green
    $response | ConvertTo-Json
} catch {
    Write-Host "✗ Root endpoint failed: $($_.Exception.Message)" -ForegroundColor Red
}

# Test health endpoint
Write-Host "`n2. Testing health endpoint..." -ForegroundColor Yellow
try {
    $response = Invoke-RestMethod -Uri "http://localhost:8080/api/v1/health" -Method GET
    Write-Host "✓ Health endpoint working" -ForegroundColor Green
    $response | ConvertTo-Json
} catch {
    Write-Host "✗ Health endpoint failed: $($_.Exception.Message)" -ForegroundColor Red
}

# Test get all users endpoint
Write-Host "`n3. Testing get all users..." -ForegroundColor Yellow
try {
    $response = Invoke-RestMethod -Uri "http://localhost:8080/api/v1/users" -Method GET
    Write-Host "✓ Get users endpoint working" -ForegroundColor Green
    $response | ConvertTo-Json
} catch {
    Write-Host "✗ Get users endpoint failed: $($_.Exception.Message)" -ForegroundColor Red
}

Write-Host "`nAPI testing completed!" -ForegroundColor Green
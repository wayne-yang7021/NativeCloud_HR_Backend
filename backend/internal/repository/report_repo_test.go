package repository

import (
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/4040www/NativeCloud_HR/internal/db"
	"github.com/4040www/NativeCloud_HR/internal/model"
	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func setupMockDB(t *testing.T) (*gorm.DB, sqlmock.Sqlmock) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("sqlmock.New() failed: %v", err)
	}
	dialector := postgres.New(postgres.Config{
		Conn: mockDB,
	})
	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		t.Fatalf("gorm.Open() failed: %v", err)
	}
	return db, mock
}

func TestGetEmployeeByID(t *testing.T) {
	dbMock, mock := setupMockDB(t)
	original := db.DB
	db.DB = dbMock
	defer func() { db.DB = original }()

	t.Run("valid employee", func(t *testing.T) {
		mock.ExpectQuery(`SELECT \* FROM "employee" WHERE employee_id = \$1.*`).
			WithArgs("emp001", 1).
			WillReturnRows(sqlmock.NewRows([]string{"employee_id", "first_name", "last_name"}).
				AddRow("emp001", "John", "Doe"))

		got, err := GetEmployeeByID("emp001")
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if got == nil || got.EmployeeID != "emp001" || got.FirstName != "John" || got.LastName != "Doe" {
			t.Errorf("unexpected result: got %+v", got)
		}
	})

	t.Run("employee not found", func(t *testing.T) {
		mock.ExpectQuery(`SELECT \* FROM "employee" WHERE employee_id = \$1.*`).
			WithArgs("emp999", 1).
			WillReturnError(gorm.ErrRecordNotFound)

		got, err := GetEmployeeByID("emp999")
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			t.Errorf("expected gorm.ErrRecordNotFound, got: %v", err)
		}
		if got != nil && got.EmployeeID != "" {
			t.Errorf("expected nil or empty result, got: %+v", got)
		}
	})

	t.Run("db error", func(t *testing.T) {
		mock.ExpectQuery(`SELECT \* FROM "employee" WHERE employee_id = \$1.*`).
			WithArgs("emp500", 1).
			WillReturnError(errors.New("db error"))

		got, err := GetEmployeeByID("emp500")
		if err == nil {
			t.Errorf("expected error, got nil")
		}
		if got == nil {
			t.Errorf("expected non-nil pointer (even on error), got nil")
		}
	})

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("unfulfilled expectations: %v", err)
	}
}

func TestGetAccessLogsByEmployeeBetween(t *testing.T) {
	dbMock, mock := setupMockDB(t)
	db.DB = dbMock

	employeeID := "emp-001"
	start := time.Date(2025, 5, 24, 0, 0, 0, 0, time.UTC)
	end := time.Date(2025, 5, 25, 0, 0, 0, 0, time.UTC)

	// mock rows
	rows := sqlmock.NewRows([]string{
		"access_id", "employee_id", "access_time", "direction", "gate_type", "gate_name", "access_result",
	}).AddRow(
		"access-001", employeeID, time.Date(2025, 5, 24, 9, 0, 0, 0, time.UTC),
		"IN", "TypeA", "Gate1", "Success",
	).AddRow(
		"access-002", employeeID, time.Date(2025, 5, 24, 18, 0, 0, 0, time.UTC),
		"OUT", "TypeB", "Gate2", "Success",
	)

	// gorm generated query 使用 Table() 要注意 regex 格式
	mock.ExpectQuery(`SELECT \* FROM "access_log" WHERE employee_id = \$1 AND access_time BETWEEN \$2 AND \$3 ORDER BY access_time asc`).
		WithArgs(employeeID, start, end).
		WillReturnRows(rows)

	// 執行
	got, err := GetAccessLogsByEmployeeBetween(employeeID, start, end)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	want := []model.AccessLog{
		{
			AccessID:     "access-001",
			EmployeeID:   employeeID,
			AccessTime:   time.Date(2025, 5, 24, 9, 0, 0, 0, time.UTC),
			Direction:    "IN",
			GateType:     "TypeA",
			GateName:     "Gate1",
			AccessResult: "Success",
		},
		{
			AccessID:     "access-002",
			EmployeeID:   employeeID,
			AccessTime:   time.Date(2025, 5, 24, 18, 0, 0, 0, time.UTC),
			Direction:    "OUT",
			GateType:     "TypeB",
			GateName:     "Gate2",
			AccessResult: "Success",
		},
	}

	if len(got) != len(want) {
		t.Fatalf("expected %d results, got %d", len(want), len(got))
	}

	for i := range got {
		if got[i] != want[i] {
			t.Errorf("mismatch at index %d: got %+v, want %+v", i, got[i], want[i])
		}
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("unfulfilled expectations: %v", err)
	}
}

func TestGetAllEmployees(t *testing.T) {
	dbMock, mock := setupMockDB(t)
	db.DB = dbMock

	// Step 3: 模擬預期的資料
	mock.ExpectQuery(`SELECT \* FROM "employee"`).
		WillReturnRows(sqlmock.NewRows([]string{
			"employee_id", "first_name", "last_name", "is_manager", "password", "email", "organization_id",
		}).AddRow(
			"E001", "John", "Doe", false, "hashedpwd", "john@example.com", "D001",
		).AddRow(
			"E002", "Jane", "Smith", true, "hashedpwd2", "jane@example.com", "D002",
		))

	want := []model.Employee{
		{
			EmployeeID:     "E001",
			FirstName:      "John",
			LastName:       "Doe",
			IsManager:      false,
			Password:       "hashedpwd",
			Email:          "john@example.com",
			OrganizationID: "D001",
		},
		{
			EmployeeID:     "E002",
			FirstName:      "Jane",
			LastName:       "Smith",
			IsManager:      true,
			Password:       "hashedpwd2",
			Email:          "jane@example.com",
			OrganizationID: "D002",
		},
	}

	// Step 4: 呼叫不接受參數的函數
	got, err := GetAllEmployees()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("GetAllEmployees() = %v, want %v", got, want)
	}

	// Step 5: 確認 mock 是否都被滿足
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %v", err)
	}
}

func TestGetDepartmentsByManager(t *testing.T) {
	dbMock, mock := setupMockDB(t)
	db.DB = dbMock

	mock.ExpectQuery(`SELECT .*organization_id FROM "employee" WHERE employee_id = \$1`).
		WithArgs("U001").
		WillReturnRows(sqlmock.NewRows([]string{"organization_id"}).
			AddRow("HR").
			AddRow("Engineering"))

	result, err := GetDepartmentsByManager("U001")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	expected := []string{"HR", "Engineering"}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}

	mock.ExpectQuery(`SELECT .*organization_id FROM "employee" WHERE employee_id = \$1`).
		WithArgs("U002").
		WillReturnRows(sqlmock.NewRows([]string{"organization_id"})) // 空查詢

	result, err = GetDepartmentsByManager("U002")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	expected = []string{}
	if result == nil {
		result = []string{}
	}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("mock expectations not met: %v", err)
	}

	mock.ExpectQuery(`SELECT .*organization_id FROM "employee" WHERE employee_id = \$1`).
		WithArgs("U003").
		WillReturnError(errors.New("mock db error"))

	_, err = GetDepartmentsByManager("U003")
	if err == nil {
		t.Errorf("expected error but got none")
	}
}

func TestGetEmployeesByDepartment(t *testing.T) {
	tests := []struct {
		name       string
		department string
		setup      func(sqlmock.Sqlmock)
		want       []model.Employee
		wantErr    bool
	}{
		{
			name:       "部門有員工",
			department: "Engineering",
			setup: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(`SELECT \* FROM "employee" WHERE organization_id = \$1`).
					WithArgs("Engineering").
					WillReturnRows(sqlmock.NewRows([]string{
						"employee_id", "first_name", "last_name", "organization_id"}).
						AddRow("E1", "Alice", "Wang", "Engineering").
						AddRow("E2", "Bob", "Chen", "Engineering"))
			},
			want: []model.Employee{
				{EmployeeID: "E1", FirstName: "Alice", LastName: "Wang", OrganizationID: "Engineering"},
				{EmployeeID: "E2", FirstName: "Bob", LastName: "Chen", OrganizationID: "Engineering"},
			},
			wantErr: false,
		},
		{
			name:       "部門無員工",
			department: "HR",
			setup: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(`SELECT \* FROM "employee" WHERE organization_id = \$1`).
					WithArgs("HR").
					WillReturnRows(sqlmock.NewRows([]string{
						"employee_id", "first_name", "last_name", "organization_id"})) // 空
			},
			want:    []model.Employee{},
			wantErr: false,
		},
		{
			name:       "資料庫錯誤",
			department: "Finance",
			setup: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(`SELECT \* FROM "employee" WHERE organization_id = \$1`).
					WithArgs("Finance").
					WillReturnError(errors.New("mock db error"))
			},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockDB, mock := setupMockDB(t)
			db.DB = mockDB
			tt.setup(mock)

			got, err := GetEmployeesByDepartment(tt.department)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetEmployeesByDepartment() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetEmployeesByDepartment() = %v, want %v", got, tt.want)
			}
		})
	}
}

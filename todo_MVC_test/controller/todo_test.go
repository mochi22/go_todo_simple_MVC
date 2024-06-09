// controller/todo_test.go
package controller

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"todo/model"
)

func TestAddTodoTableDriven(t *testing.T) {
	testDatas := map[string]struct {
		name           string
		form           string
		expectedStatus int
		expectedTodos  int
		expectedItem   string
	}{
		"simple space": {
			name:           "Add valid todo",
			form:           "item=TestTodo",
			expectedStatus: http.StatusFound,
			expectedTodos:  1,
			expectedItem:   "TestTodo",
		},
		"test space and any special charactor": {
			name:           "Add valid todo",
			form:           "item=Test%20%21%23%24%25%26Todo",
			expectedStatus: http.StatusFound,
			expectedTodos:  2,
			expectedItem:   "Test !#$%&Todo",
		},
	}

	for _, tt := range testDatas {

		t.Run(tt.name, func(t *testing.T) {
			// テスト用のリクエストを作成
			req, err := http.NewRequest("POST", "/todo", strings.NewReader(tt.form))
			if err != nil {
				t.Fatal(err)
			}
			req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

			// レスポンスを記録するためのレスポンスレコーダーを作成
			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(AddTodo)

			// ハンドラを呼び出す
			handler.ServeHTTP(rr, req)

			// ステータスコードが期待されるものであることを確認
			if status := rr.Code; status != tt.expectedStatus {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, tt.expectedStatus)
			}

			// Todo が正しく追加されたことを確認
			if len(todos) != tt.expectedTodos || todos[len(todos)-1].Item != tt.expectedItem {
				t.Errorf("todo was not added correctly: got %v", todos)
			}
		})
	}
}

func TestUpdateTodoTableDriven(t *testing.T) {
	// テスト用のデータをセットアップ
	todos = []model.Todo{
		{ID: 1, Item: "Test Todo", Completed: false},
	}

	testDatas := []struct {
		name           string
		form           string
		expectedStatus int
		expectedItem   string
		expectedComp   bool
	}{
		{
			name:           "Update existing todo",
			form:           "id=1&item=Updated%20Todo&completed=true",
			expectedStatus: http.StatusFound,
			expectedItem:   "Updated Todo",
			expectedComp:   true,
		},
	}

	for _, tt := range testDatas {
		t.Run(tt.name, func(t *testing.T) {
			// テスト用のリクエストを作成
			req, err := http.NewRequest("POST", "/todo/update", strings.NewReader(tt.form))
			if err != nil {
				t.Fatal(err)
			}
			req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

			// レスポンスを記録するためのレスポンスレコーダーを作成
			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(UpdateTodo)

			// ハンドラを呼び出す
			handler.ServeHTTP(rr, req)

			// ステータスコードが期待されるものであることを確認
			if status := rr.Code; status != tt.expectedStatus {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, tt.expectedStatus)
			}

			// Todo が正しく更新されたことを確認
			if todos[0].Item != tt.expectedItem || todos[0].Completed != tt.expectedComp {
				t.Errorf("todo was not updated correctly: got %v", todos)
			}
		})
	}
}

func TestDeleteTodoTableDriven(t *testing.T) {
	testDatas := []struct {
		name           string
		form           string
		setupTodos     []model.Todo
		expectedStatus int
		expectedTodos  int
	}{
		{
			name:           "Delete existing todo",
			form:           "id=1",
			setupTodos:     []model.Todo{{ID: 1, Item: "Test Todo", Completed: false}},
			expectedStatus: http.StatusFound,
			expectedTodos:  0,
		},
	}

	for _, tt := range testDatas {
		t.Run(tt.name, func(t *testing.T) {
			// テスト用のデータをセットアップ
			todos = tt.setupTodos

			// テスト用のリクエストを作成
			req, err := http.NewRequest("POST", "/todo/delete", strings.NewReader(tt.form))
			if err != nil {
				t.Fatal(err)
			}
			req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

			// レスポンスを記録するためのレスポンスレコーダーを作成
			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(DeleteTodo)

			// ハンドラを呼び出す
			handler.ServeHTTP(rr, req)

			// ステータスコードが期待されるものであることを確認
			if status := rr.Code; status != tt.expectedStatus {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, tt.expectedStatus)
			}

			// Todo が正しく削除されたことを確認
			if len(todos) != tt.expectedTodos {
				t.Errorf("todo was not deleted correctly: got %v", todos)
			}
		})
	}
}

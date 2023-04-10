import requests
import unittest
import random

url = "http://127.0.0.1:8080/todos"

todo = {
    "id": str(-random.randint(0, 999)),
    "title": "todo",
    "description": "description",
    "complete": False,
}


def test_get_todos():
    req = requests.get(url)
    data = req.json()
    assert data[0]["id"] != None


def test_post_todo():
    req = requests.post(url, data=todo)
    assert req.status_code == 200


def test_get_todo_by_id():
    req = requests.get(f"{url}/{todo['id']}")
    data = req.json()
    assert data["id"] == todo["id"]

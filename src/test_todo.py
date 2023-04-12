import requests
import random
from dotenv import load_dotenv
import os

load_dotenv()

url = f'http://{os.getenv("APIHOST")}:{os.getenv("APIPORT")}/todos'


test_user = {
    "id": 1,
    "name": "User",
    "password": "1234",
    "mail": "user@mail.com",
    "session_key": "123",
}

todo = {
    "id": str(-random.randint(0, 999)),
    "title": "todo",
    "description": "description",
    "complete": False,
}


def test_get_todos():
    req = requests.get(url)
    assert req.status_code == 200
    data = req.json()
    assert data[0]["user_id"] != None


def test_post_todo():
    req = requests.post(url, data=todo, headers=prepare_headers(test_user))
    assert req.status_code == 200


def test_get_todo_by_id():
    req = requests.get(f"{url}/{todo['id']}", headers=prepare_headers(test_user))
    assert req.status_code == 200
    data = req.json()
    assert data["user_id"] == todo["user_id"]


def prepare_headers(user: object) -> object:
    return {"User_id": str(user["id"]), "Session_key": str(user["session_key"])}

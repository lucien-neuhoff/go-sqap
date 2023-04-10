import requests
import random
from dotenv import load_dotenv
import os

load_dotenv()

url = f'http://{os.getenv("APIHOST")}:{os.getenv("APIPORT")}/todos'

todo = {
    "id": str(-random.randint(0, 999)),
    # Todo: get user id after user creation
    "user_id": ""
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
    req = requests.post(url, data=todo)
    assert req.status_code == 200


def test_get_todo_by_id():
    req = requests.get(f"{url}/{todo['user_id']}/{todo[]}")
    assert req.status_code == 200
    data = req.json()
    assert data["user_id"] == todo["user_id"]

import requests
import random
from werkzeug.security import generate_password_hash
from dotenv import load_dotenv
import os

load_dotenv()

url = f'http://{os.getenv("APIHOST")}:{os.getenv("APIPORT")}/users/auth'

rand = random.randint(100, 999)
user = {
    "name": f"User{rand}",
    "password": generate_password_hash("password1234", "sha256"),
    "email": f"user{rand}@mail.com",
}


def test_auth_user_signup():
    req = requests.post(f"{url}/signup", data=user)
    assert req.status_code == 200


def test_auth_user_signin():
    req = requests.get(f"{url}/signin/{user['email']}/{user['password']}")
    assert req.status_code == 200
    data = req.json()
    assert data["email"] == user["email"]

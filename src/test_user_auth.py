import requests
import random
from werkzeug.security import generate_password_hash
from dotenv import load_dotenv
import os

load_dotenv()

url = f'http://{os.getenv("APIHOST")}:{os.getenv("APIPORT")}/users/auth'

rand = random.randint(100, 999)
user = {
    "username": f"User{rand}",
    "password": "1234",
    "email": f"user{rand}@mail.com",
}

print()


def test_auth_user_signup():
    req = requests.post(f"{url}/signup", data=user)
    assert req.status_code == 200


def test_auth_user_signup_no_email():
    u = user.copy()
    del u["email"]
    req = requests.post(f"{url}/signup", data=u)
    data = req.json()
    assert req.status_code == 400 and data["message"] == "auth/missing-email"


def test_auth_user_signup_no_password():
    u = user.copy()
    del u["password"]
    req = requests.post(f"{url}/signup", data=u)
    data = req.json()
    assert req.status_code == 400 and data["message"] == "auth/missing-password"


def test_auth_user_signup_no_username():
    u = user.copy()
    del u["username"]
    req = requests.post(f"{url}/signup", data=u)
    data = req.json()
    assert req.status_code == 400 and data["message"] == "auth/missing-username"


def test_auth_user_signin():
    req = requests.get(f"{url}/signin/{user['email']}/{user['password']}")
    assert req.status_code == 200
    data = req.json()
    assert data["email"] == user["email"]

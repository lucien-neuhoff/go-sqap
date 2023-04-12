### User

-   <u>Hash</u> the **password** on the _client_ end
-   <u>**TLS**</u>
-   #### Session system
    -   <u>Create</u> a **session_key** on _user_ auth
    -   <u>Save</u> the **session_start_at** on _user_ auth
    -   <u>Send</u> the **session_key** to the _user_
    -   <u>Disconnect</u> the _user_ when **session_start_at** exceeds a set amount of time
    -   Any _user_ action should pass the **user_id** through the **header** as **User_id**

---

### Auth

-   **Error messages**:
    -   `auth/missing-email` No email were given
    -   `auth/missing-password` No password were given
    -   `auth/missing-username` No username were given
    -   `auth/email-not-found` No user was found with the given email
    -   `auth/password-mismatch` The given password does not match the one associated with the given email

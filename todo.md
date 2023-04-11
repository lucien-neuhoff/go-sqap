### User

-   <u>Hash</u> the **password** on the _client_ end
-   #### Session system

    -   <u>Create</u> a **session_key** on _user_ auth
    -   <u>Save</u> the **session_start_at** on _user_ auth
    -   <u>Send</u> the **session_key** to the _user_
    -   <u>Disconnect</u> the _user_ when **session_start_at** exceeds a set amount of time

---

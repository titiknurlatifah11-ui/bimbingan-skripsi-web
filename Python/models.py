class User:
    def __init__(self, username, password, role):
        self.username = username
        self.password = password
        self.role = role

    def login(self, username, password):
        return self.username == username and self.password == password


class Admin(User):
    def __init__(self, username, password):
        super().__init__(username, password, 'admin')

    def manage_school(self):
        return "Admin mengelola sekolah."


class Student(User):
    def __init__(self, username, password):
        super().__init__(username, password, 'siswa')

    def view_grades(self):
        return "Siswa melihat nilai."


class Teacher(User):
    def __init__(self, username, password):
        super().__init__(username, password, 'guru')

    def manage_classes(self):
        return "Guru mengelola kelas."


class SchoolSystem:
    def __init__(self):
        self.users = []

    def add_user(self, user):
        self.users.append(user)

    def authenticate(self, username, password):
        for user in self.users:
            if user.login(username, password):
                return user
        return None

from models import Admin, Student, Teacher, SchoolSystem

def main():
    school_system = SchoolSystem()

    # Tambahkan beberapa pengguna
    school_system.add_user(Admin('admin', 'admin123'))
    school_system.add_user(Student('siswa', 'siswa123'))
    school_system.add_user(Teacher('guru', 'guru123'))

    print("Selamat datang di Sistem Administrasi Sekolah")
    username = input("Masukkan username: ")
    password = input("Masukkan password: ")

    user = school_system.authenticate(username, password)
    if user:
        print(f"Login berhasil sebagai {user.role}")
        if user.role == 'admin':
            print(user.manage_school())
        elif user.role == 'siswa':
            print(user.view_grades())
        elif user.role == 'guru':
            print(user.manage_classes())
    else:
        print("Login gagal, username atau password salah.")

if __name__ == "__main__":
    main()

def sorting_algorithm(A):
    n = len(A)
    for i in range(n):
        swapped = False
        for j in range(n - i - 1):
            if A[j] > A[j + 1]:
                A[j], A[j + 1] = A[j + 1], A[j]
                swapped = True
        if not swapped:
            break

arr = [30, 40, 20, 10, 50, 60, 70]
sorting_algorithm(arr)
print("Array setelah diurutkan:")
print(arr)
# by Don24Kane    

# Python

def merge(v, left, mid, right):
    n1 = (len(v) // 2) + 1
    n2 = len(v) // 2

    L = v[left:left + n1]
    R = v[mid + 1:mid + 1 + n2]

    i = 0
    j = 0
    k = left

    while i < n1 and j < n2:
        if L[i] <= R[j]:
            v[k] = L[i]
            i += 1
        else:
            v[k] = R[j]
            j += 1
        k += 1

    while i < n1:
        v[k] = L[i]
        i += 1
        k += 1

    while j < n2:
        v[k] = R[j]
        j += 1
        k += 1

def mergeSort(v, left, right):
    if left >= right:
        return

    mid = left + (right - left) // 2
    mergeSort(v, left, mid)
    mergeSort(v, mid + 1, right)
    merge(v, left, mid, right)

def printSolution(v):
    for i in range(len(v)):
        print(v[i], end=" ")
    print()

if __name__ == "__main__":
    v = []
    n = int(input("Number of elements: "))
    print("Elements: ", end="")
    v = [int(input()) for _ in range(n)]

    mergeSort(v, 0, n - 1)

    print("Sorted vector: ")
    printSolution(v)


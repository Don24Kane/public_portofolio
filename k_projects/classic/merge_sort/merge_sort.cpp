# by Don24Kane

# C++

#include <bits/stdc++.h>
using namespace std;

void merge(vector<int>& v, int left, int mid, int right)
{
    int n1 = (v.size() / 2) + 1;
    int n2 = v.size() / 2;

    vector<int> L(n1), R(n2);

    for (auto i = 0; i < n1; i++)
        L[i] = v[left + i];
    for (auto j = 0; j < n2; j++)
        R[j] = v[mid + 1 + j];

    int i = 0, j = 0;
    int k = left;

    while (i < n1 && j < n2) {
        if (L[i] <= R[j]) {
            v[k] = L[i];
            i++;
        }
        else {
            v[k] = R[j];
            j++;
        }
        k++;
    }

    while (i < n1) {
        v[k] = L[i];
        i++;
        k++;
    }

    while (j < n2) {
        v[k] = R[j];
        j++;
        k++;
    }
}

void mergeSort(vector<int>& v, int left, int right)
{
    if (left >= right)
        return;

    int mid = left + (right - left) / 2;
    mergeSort(v, left, mid);
    mergeSort(v, mid + 1, right);
    merge(v, left, mid, right);
}

void printSolution(vector<int>& v)
{
    for (auto i = 0; i < v.size(); i++)
        cout << v[i] << " ";
    cout << endl;
}

int main()
{
    vector<int> v = { };
    int n;
    cout << "Number of elements : "; cin >> n; cout << endl;
    cout << "Elements : "; for (auto i = 0; i <= n; i++) cin >> v[i]; cout << endl;

    mergeSort(v, 0, n - 1);

    cout << "Sorted vector : \n";
    printSolution(v);
    return 0;
}

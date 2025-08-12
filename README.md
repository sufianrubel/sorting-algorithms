# Sorting Algorithms

A collection of common sorting algorithm implementations in **Golang**, including:

- Quick Sort
- Merge Sort
- Heap Sort
- Bubble Sort
- Insertion Sort
- Selection Sort
- Radix Sort

## 📚 Overview

This repository is intended for **Data Structures and Algorithms (DSA)** practice using **Go**.  
Each algorithm is implemented in a separate file under the `algorithms/` folder, with clear and simple code.

---

## 📊 Time Complexities

| Algorithm       | Best Time   | Average Time | Worst Time  | Space Complexity | Stable |
|-----------------|-------------|--------------|--------------|------------------|--------|
| Quick Sort      | O(n log n)  | O(n log n)   | O(n²)        | O(log n)         | No     |
| Merge Sort      | O(n log n)  | O(n log n)   | O(n log n)   | O(n)             | Yes    |
| Heap Sort       | O(n log n)  | O(n log n)   | O(n log n)   | O(1)             | No     |
| Bubble Sort     | O(n)        | O(n²)        | O(n²)        | O(1)             | Yes    |
| Insertion Sort  | O(n)        | O(n²)        | O(n²)        | O(1)             | Yes    |
| Selection Sort  | O(n²)       | O(n²)        | O(n²)        | O(1)             | No     |
| Radix Sort      | O(nk)       | O(nk)        | O(nk)        | O(n + k)         | Yes    |

---

## 🚀 Running the Project

1. **Clone the repository**
```bash
git clone https://github.com/sufianrubel/sorting-algorithms.git
cd sorting-algorithms
 ```
2. **Initialize Go module**
```bash
go mod init sorting-algorithms
 ```
3. **Run the program**
```bash
go run main.go
 ```

package adminService

import (
	"fmt"
	"math"
	"sort"
)

// Alternative merepresentasikan data awal untuk setiap alternatif
type Alternative struct {
	Name   string
	Values []float64
}

// RankedAlternative digunakan untuk menyimpan hasil akhir beserta skornya
type RankedAlternative struct {
	Name  string
	Score float64
	Rank  int
}

func main() {
	// --- DATA INPUT ---
	// Masukkan data alternatif Anda di sini
	alternatives := []Alternative{
		{Name: "Alt1", Values: []float64{9, 20, 7, 3, 15000, 9, 4.8}},
		{Name: "Alt2", Values: []float64{8, 10, 9, 5, 5000, 8, 4.5}},
		{Name: "Alt3", Values: []float64{7, 25, 6, 4, 10000, 8, 4.2}},
		{Name: "Alt4", Values: []float64{10, 30, 8, 4, 25000, 10, 4.9}},
		{Name: "Alt5", Values: []float64{9, 15, 9, 5, 20000, 9, 4.7}},
		{Name: "Alt6", Values: []float64{8, 18, 8, 3, 10000, 9, 4.4}},
		{Name: "Alt7", Values: []float64{7, 22, 7, 4, 12000, 8, 4.6}},
		{Name: "Alt8", Values: []float64{8, 28, 7, 3, 30000, 9, 4.5}},
	}

	// Masukkan bobot untuk setiap kriteria (C1 sampai C7)
	// Pastikan jumlah bobot sama dengan jumlah kriteria
	weights := []float64{0.20, 0.15, 0.15, 0.10, 0.10, 0.15, 0.15}

	// Tentukan jenis kriteria: "benefit" atau "cost"
	criteriaTypes := []string{"benefit", "cost", "benefit", "cost", "benefit", "benefit", "benefit"}

	// --- PROSES PERHITUNGAN MABAC ---

	// Langkah 1: Normalisasi Matriks Keputusan (N)
	normalizedMatrix := normalize(alternatives, criteriaTypes)

	// Langkah 2: Menghitung Matriks Tertimbang (V)
	weightedMatrix := calculateWeightedMatrix(normalizedMatrix, weights)

	// Langkah 3: Menentukan Matriks Perkiraan Batas (G)
	borderApproximationArea := calculateBorderApproximationArea(weightedMatrix)

	// Langkah 4: Menghitung Jarak Alternatif dari Batas (Q)
	distanceMatrix := calculateDistanceMatrix(weightedMatrix, borderApproximationArea)

	// Langkah 5: Menghitung Skor Akhir dan Melakukan Perankingan (S)
	finalRanking := calculateFinalScores(distanceMatrix, alternatives)

	// --- TAMPILKAN HASIL ---
	fmt.Println("Hasil Akhir Perankingan Metode MABAC:")
	fmt.Println("-------------------------------------------")
	fmt.Printf("%-5s %-15s %-15s\n", "Rank", "Alternative", "Score (S)")
	fmt.Println("-------------------------------------------")
	for _, result := range finalRanking {
		fmt.Printf("%-5d %-15s %-15.4f\n", result.Rank, result.Name, result.Score)
	}
	fmt.Println("-------------------------------------------")
}

// normalize melakukan normalisasi matriks keputusan
// Rumus Benefit: (Xij - Xi_min) / (Xi_max - Xi_min)
// Rumus Cost:   (Xi_max - Xij) / (Xi_max - Xi_min)
func normalize(alternatives []Alternative, criteriaTypes []string) [][]float64 {
	numAlternatives := len(alternatives)
	numCriteria := len(criteriaTypes)
	normalized := make([][]float64, numAlternatives)

	for i := 0; i < numCriteria; i++ {
		minVal := alternatives[0].Values[i]
		maxVal := alternatives[0].Values[i]
		for j := 1; j < numAlternatives; j++ {
			if alternatives[j].Values[i] < minVal {
				minVal = alternatives[j].Values[i]
			}
			if alternatives[j].Values[i] > maxVal {
				maxVal = alternatives[j].Values[i]
			}
		}

		for j := 0; j < numAlternatives; j++ {
			if i == 0 {
				normalized[j] = make([]float64, numCriteria)
			}
			val := alternatives[j].Values[i]
			rangeVal := maxVal - minVal
			if rangeVal == 0 {
				normalized[j][i] = 1 // Jika semua nilai sama
				continue
			}

			if criteriaTypes[i] == "benefit" {
				normalized[j][i] = (val - minVal) / rangeVal
			} else if criteriaTypes[i] == "cost" {
				normalized[j][i] = (maxVal - val) / rangeVal
			}
		}
	}
	return normalized
}

// calculateWeightedMatrix menghitung matriks tertimbang (V)
// Rumus: Vij = (wi * Nij) + wi
func calculateWeightedMatrix(normalizedMatrix [][]float64, weights []float64) [][]float64 {
	numRows := len(normalizedMatrix)
	numCols := len(weights)
	weighted := make([][]float64, numRows)

	for i := 0; i < numRows; i++ {
		weighted[i] = make([]float64, numCols)
		for j := 0; j < numCols; j++ {
			weighted[i][j] = (weights[j] * normalizedMatrix[i][j]) + weights[j]
		}
	}
	return weighted
}

// calculateBorderApproximationArea menghitung matriks perkiraan batas (G)
// Rumus: Gi = (Product(Vij for j=1 to m))^(1/m) - Rata-rata Geometris
func calculateBorderApproximationArea(weightedMatrix [][]float64) []float64 {
	numRows := len(weightedMatrix)
	if numRows == 0 {
		return []float64{}
	}
	numCols := len(weightedMatrix[0])
	g := make([]float64, numCols)

	for j := 0; j < numCols; j++ {
		product := 1.0
		for i := 0; i < numRows; i++ {
			product *= weightedMatrix[i][j]
		}
		g[j] = math.Pow(product, 1.0/float64(numRows))
	}
	return g
}

// calculateDistanceMatrix menghitung matriks jarak dari batas (Q)
// Rumus: Qij = Vij - Gi
func calculateDistanceMatrix(weightedMatrix [][]float64, g []float64) [][]float64 {
	numRows := len(weightedMatrix)
	numCols := len(g)
	q := make([][]float64, numRows)

	for i := 0; i < numRows; i++ {
		q[i] = make([]float64, numCols)
		for j := 0; j < numCols; j++ {
			q[i][j] = weightedMatrix[i][j] - g[j]
		}
	}
	return q
}

// calculateFinalScores menghitung skor akhir (S) dan melakukan perankingan
// Rumus: Sj = Sum(Qij for i=1 to n)
func calculateFinalScores(distanceMatrix [][]float64, alternatives []Alternative) []RankedAlternative {
	numAlternatives := len(distanceMatrix)
	finalScores := make([]RankedAlternative, numAlternatives)

	for i := 0; i < numAlternatives; i++ {
		sum := 0.0
		for _, qVal := range distanceMatrix[i] {
			sum += qVal
		}
		finalScores[i] = RankedAlternative{
			Name:  alternatives[i].Name,
			Score: sum,
		}
	}

	// Lakukan sorting dari skor tertinggi ke terendah
	sort.Slice(finalScores, func(i, j int) bool {
		return finalScores[i].Score > finalScores[j].Score
	})

	// Tetapkan peringkat
	for i := range finalScores {
		finalScores[i].Rank = i + 1
	}

	return finalScores
}
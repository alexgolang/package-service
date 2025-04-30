package service

import (
	"context"
	"math"
)

type Logger interface {
	Println(v ...interface{})
	Printf(format string, v ...interface{})
}

type PackageService struct {
	packageSizes []int
	logger Logger
}

func NewPackageService(packageSizes []int, logger Logger) *PackageService {
	return &PackageService{packageSizes: packageSizes, logger: logger}
}

func (s *PackageService) GetPackageSize(ctx context.Context, target int) map[int]int {
	s.logger.Printf("GetPackageSize %d", target)
	result := getSize(target, s.packageSizes)
	s.logger.Printf("result for %d is %v", target, result)
	return result
}

func getSize(target int, packSizes []int) map[int]int {
    // dp[i] will store the minimum total we can make >= i
    dp := make([]int, target + packSizes[0]) // to handle overflow
    prevPack := make([]int, target + packSizes[0]) // to reconstruct the solution
    
    // Initialize with max value
    for i := range dp {
        dp[i] = math.MaxInt32
    }
    dp[0] = 0
    
    // For each amount up to target
    for i := 0; i <= target; i++ {
        // If this amount can be made
        if dp[i] != math.MaxInt32 {
            // Try adding each pack size
            for _, pack := range packSizes {
                if i + pack < len(dp) && dp[i] + 1 < dp[i + pack] {
                    dp[i + pack] = dp[i] + 1
                    prevPack[i + pack] = pack
                }
            }
        }
    }
    
    // Find the smallest total >= target
    minTotal := target
    for minTotal < len(dp) && dp[minTotal] == math.MaxInt32 {
        minTotal++
    }
    
    // Reconstruct the solution
    result := make(map[int]int)
    current := minTotal
    for current > 0 {
        pack := prevPack[current]
        result[pack]++
        current -= pack
    }
    
    return result
}


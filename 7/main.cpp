#include <iostream>
#include <fstream>
#include <utility>
#include <vector>
#include <string>
#include <sstream>
#include <cmath>

const std::string FILE_NAME = "input.txt";

// Read input from file
// Expected input: <value>: <number> .. <number>
std::vector<std::pair<long, std::vector<int>>> getEquations() {
    std::ifstream file(FILE_NAME);
    if (!file.is_open()) {
        exit(-1);
    }

    std::vector<std::pair<long, std::vector<int>>> equations;
    std::string line{};
    while (std::getline(file, line)) {
        size_t delim = line.find(":");
        if (delim == std::string::npos) {
            exit(-1);
        }

        long val = std::stoll(line.substr(0, delim));
        std::vector<int> numbers;
        std::istringstream iss(line.substr(delim + 1));
        int tmp;
        while (iss >> tmp) {
            numbers.push_back(tmp);
        }

        equations.push_back({val, numbers});
    }
    file.close();
    return equations;
}

// Print equation value and numbers
void printEquations(std::vector<std::pair<long, std::vector<int>>> equations) {
    for (const auto& [val, numbers] : equations) {
        std::cout << "Test val: " << val << " Numbers: ";
        for (const auto& number : numbers) std::cout << number << ' ';
        std::cout << '\n';
    }
}

// Perform one set of combinations on numbers
void performOperations(int& operation, long& result, std::vector<int> numbers) {
    for (int i = 1; i < numbers.size(); ++i) {
        char op = (operation % 2 == 0) ? '+' : '*';
        operation /= 2;

        if (op == '+') {
            result += numbers[i];
        } else if (op == '*') {
            result *= numbers[i];
        }
    }
}

long checkCombinations(const std::vector<int> numbers, long target, int total_combination) {
    for (int combination = 0; combination < total_combination; ++combination) {
        long result = numbers[0];
        int curr_combination = combination;
        performOperations(curr_combination, result, numbers);
        if (result == target) {
            return target;
        }
    }
    return 0;
}

// Return value of equation if valid for part 1, otherwise return 0.
long getEquationSumPart1(const std::pair<long, std::vector<int>> equation) {
    long target = equation.first;
    std::vector<int> numbers = equation.second;

    int num_operators = numbers.size() - 1;
    int total_combination = std::pow(2, num_operators);

    return checkCombinations(numbers, target, total_combination);
}

// Return value of equation if valid for part 2, otherwise return 0.
long getEquationSumPart2(const std::pair<long, std::vector<int>> equation) {
    long target = equation.first;
    std::vector<int> numbers = equation.second;
    int num_operators = numbers.size() - 1;
    int total_combination = std::pow(2, num_operators);

    long result = checkCombinations(numbers, target, total_combination);
    if (result != 0) {
        return result;
    }

    // Check if equation is valid if using '|'
    for (int i = 0; i < numbers.size(); ++i) {
       long  
    }

    
    return 0;
}

// Get sum of valid calibration results
long getTotalSum(std::vector<std::pair<long, std::vector<int>>> equations, bool part1) {
    long sum = 0;

    for (const auto e : equations) {
        sum += (part1) ? getEquationSumPart1(e) : getEquationSumPart2(e);
    }

    return sum;
}

int main(void) {
    auto equations = getEquations();
    auto sum = getTotalSum(equations, true);
    std::cout << "Sum: " << sum << std::endl;
    return 0;
}

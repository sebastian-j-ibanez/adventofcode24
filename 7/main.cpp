#include <iostream>
#include <fstream>
#include <utility>
#include <vector>
#include <string>
#include <sstream>
#include <cmath>

const std::string FILE_NAME = "input.txt";

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

void printEquations(std::vector<std::pair<long, std::vector<int>>> equations) {
    for (const auto& [val, numbers] : equations) {
        std::cout << "Test val: " << val << " Numbers: ";
        for (const auto& number : numbers) std::cout << number << ' ';
        std::cout << '\n';
    }
}

//long getEquationSum(const std::pair<long, std::vector<int>> equation) {
//    long target = equation.first;
//    std::vector<int> numbers = equation.second;
//
//    int num_operators = numbers.size() - 1;
//    int total_combination = std::pow(2, num_operators);
//
//    for (int combination = 0; combination < total_combination; ++combination) {
//        long result = numbers[0];
//        int curr_combination = combination;
//
//        for (int i = 1; i < numbers.size(); ++i) {
//            char op = (curr_combination % 2 == 0) ? '+' : '*';
//            curr_combination /= 2;
//
//            if (op == '+') {
//                result += numbers[i];
//            } else if (op == '*') {
//                result *= numbers[i];
//            }
//        }
//
//        if (result == target) {
//            return target;
//        }
//    }
//    return 0;
//}


long getEquationSum(const std::pair<long, std::vector<int>> equation) {
    long target = equation.first;
    std::vector<int> numbers = equation.second;

    int num_operators = numbers.size() - 1;
    int total_combination = std::pow(2, num_operators);

    for (int combination = 0; combination < total_combination; ++combination) {
        long result = numbers[0];
        int curr_combination = combination;

        for (int i = 1; i < numbers.size(); ++i) {
            char op = (curr_combination % 2 == 0) ? '+' : '*';
            curr_combination /= 2;

            if (op == '+') {
                result += numbers[i];
            } else if (op == '*') {
                result *= numbers[i];
            }
        }

        if (result == target) {
            return target;
        }
    }
    return 0;
}

// Get sum of valid calibration results
long getTotalSum(std::vector<std::pair<long, std::vector<int>>> equations) {
    long sum = 0;

    for (const auto e : equations) {
        sum += getEquationSum(e);
    }

    return sum;
}

int main(void) {
    auto equations = getEquations();
    auto sum = getTotalSum(equations);
    std::cout << "Sum: " << sum << std::endl;
    return 0;
}

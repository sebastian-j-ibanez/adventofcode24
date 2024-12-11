#include <iostream>
#include <fstream>
#include <numeric>
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

int getEquationSum(const std::pair<long, std::vector<int>> equation) {
    long target = equation.first;
    std::vector<int> numbers = equation.second;
    std::vector<char> operators{ '+', '*'};

    for (int i = 1; i < numbers.size(); i++) {
        long result = numbers[0];

        for (int k = 1; k < numbers.size(); k++) {
            for (int j = 0; j < operators.size(); j++) {
                if (operators[j] == '+') {
                    result += numbers[k];
                } else if (operators[j] == '*') {
                    result *= numbers[k];
                }

                // else {
                //     if (operators[j] == '+') {
                //         result -= numbers[i];
                //     } else if (operators[j] == '*') {
                //         result /= numbers[i];
                //     }
                // }
            }
        }
        if (result == target) {
                    return std::accumulate(numbers.begin(), numbers.end(), 0);
        }
    }
    return 0;
}

// Get sum of valid calibration results
int getTotalSum(std::vector<std::pair<long, std::vector<int>>> equations) {
    int sum;

    for (const auto e : equations) {
        sum += getEquationSum(e);
    }

    return sum;
}

int main(void) {
    auto equations = getEquations();
    int sum = getTotalSum(equations);
    std::cout << "Sum: " << sum << std::endl;
    return 0;
}

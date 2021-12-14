#include <iostream>
#include <fstream>
#include <string>
#include <chrono>

using namespace std;

int compare_Nums(int a, int b) { //b is old num, a is new num
	if (a > b) {
		return 1;
	}
	else {
		return 0;
	}
}

int get_Increases() {
	int tmp = 0;
	fstream inputfile;
	inputfile.open("input.txt", ios::in); //in is for read, out is for write
	if (inputfile.is_open()) {
		string tp;
		int prevvar;
		int linenum;
		linenum = 0;
		while (getline(inputfile, tp)) { //read data from file and put in string
			linenum += 1;
			if (linenum > 1) {
				tmp += compare_Nums(std::stoi(tp), prevvar);
			}
			prevvar = std::stoi(tp);
		}
	}
	return tmp;
}

int get3Nums() {
	int tmp = 0;
	int prevvar1 = 0;
	int prevvar2 = 0;
	int currtotal = 0;
	int prevtotal = 0;
	fstream inputfile;
	inputfile.open("input.txt", ios::in); //in is for read, out is for write
	if (inputfile.is_open()) {
		string tp;


		int linenum;
		linenum = 0;
		while (getline(inputfile, tp)) { //read data from file and put in string
			linenum += 1;
			if (linenum > 3) {
				currtotal = std::stoi(tp) + prevvar1 + prevvar2;
				tmp += compare_Nums(currtotal, prevtotal);
			}
			prevvar2 = prevvar1;
			prevvar1 = std::stoi(tp);
			prevtotal = currtotal;
		}
	}
	return tmp;
}

int main() {

	using std::chrono::high_resolution_clock;
	using std::chrono::duration_cast;
	using std::chrono::duration;
	using std::chrono::milliseconds;

	auto t1 = high_resolution_clock::now();
	cout << "Individual comparisons";
	int total = get_Increases();
	cout << "\nTotal increases was: " << total;
	auto t2 = high_resolution_clock::now();

	/* Getting number of milliseconds as an integer. */
	auto ms_int = duration_cast<milliseconds>(t2 - t1);

	/* Getting number of milliseconds as a double. */
	duration<double, std::milli> ms_double = t2 - t1;

	std::cout << "\n" << ms_int.count() << "ms\n";
	std::cout << ms_double.count() << "ms";

	t1 = high_resolution_clock::now();
	cout << "\nTesting with batches of 3 nums";
	total = get3Nums();
	cout << "\nTotal increases was: " << total;
	t2 = high_resolution_clock::now();

	ms_int = duration_cast<milliseconds>(t2 - t1);
	duration<double, std::milli> ms_doublev2 = t2 - t1;
	std::cout << "\n" << ms_int.count() << "ms\n";
	std::cout << ms_doublev2.count() << "ms";

	return 0;

}
#include "merkletree.h"
#include <filesystem>
#include <fstream>
#include <iostream>
#include <vector>

using std::cout, std::string, std::vector, std::exit;
using std::ifstream;
using std::ofstream;
using std::string_view;
using std::filesystem::path;

constexpr string_view OK = "\033[32;1;4mOK\033[0m ";
constexpr string_view ERROR = "\033[31;1;4mERROR\033[0m ";
constexpr string_view CHECKSUM_FILE = "CHECKSUM";
constexpr string_view HELP = R"(
NAME
irow - Immutable Records using One-Way Hash

SYNOPSIS
irow generate [directory]
irow verify   [directory]

DESCRIPTION
Given a directory of files, a snapshot of its contents is built using a Merkle
hash tree. The hash of root the is stored in a file called CHECKSUM. If any of
the files in the directory, or the CHECKSUM was tampered with, the verification
will give a warning.
)";

void help() {
  cout << HELP << "\n";
  exit(1);
}

void check(bool condition, const string &message) {
  if (!condition) {
    std::cerr << ERROR << message << std::endl;
    std::exit(1);
  }
}

void generate(const path &directory, const char action) {
  auto filenames = merkletree::get_filenames(directory);
  check(not filenames.empty(), "Failed to get filenames");

  auto checksum_path = directory / CHECKSUM_FILE;
  auto blocks = merkletree::get_blocks(filenames);
  auto root_hash = merkletree::find_root(blocks);

  if (action == 'v') {
    auto is_checksum_exists = std::filesystem::exists(checksum_path);
    ifstream infile(checksum_path);
    string checksum;
    infile >> checksum;
    infile.close();
    check(is_checksum_exists, "CHECKSUM does not exist");
    check(checksum == root_hash, "Computed CHECKSUM did NOT match");

    std::cout << OK << directory << std::endl;
    return;
  }

  if (action == 'g') {
    ofstream outfile(checksum_path);
    outfile.seekp(0);
    outfile << root_hash;
    outfile.close();
    std::cout << OK << "Created CHECKSUM" << std::endl;
    return;
  }
}

int main(int argc, char *argv[]) {
  if (argc < 3)
    help();

  auto action = argv[1][0];
  path directory{argv[2]};

  if (action == 'g' or action == 'v') {
    generate(directory, action);
  } else {
    help();
  }

  return 0;
}

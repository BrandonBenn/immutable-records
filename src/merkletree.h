#pragma once

#include <filesystem>
#include <memory>
#include <optional>
#include <string>
#include <vector>

using std::optional;
using std::string;
using std::unique_ptr;
using std::vector;
using std::filesystem::path;

namespace merkletree {

string find_root(vector<string> &);
string hash(const string &content);
vector<string> get_filenames(const path &);
vector<string> get_blocks(const vector<string> &);

} // namespace merkletree

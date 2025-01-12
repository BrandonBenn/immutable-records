#include "merkletree.h"

#include <filesystem>
#include <fstream>
#include <iostream>
#include <openssl/evp.h>
#include <sstream>

using std::string;
using std::stringstream;
using std::vector;
using std::filesystem::path;
using std::filesystem::recursive_directory_iterator;

namespace merkletree {

string find_root(vector<string> &blocks) {
  auto current_level = blocks;

  while (current_level.size() > 1) {
    vector<string> next_level;

    for (auto i = 0; i < current_level.size(); i = i + 2) {
      string concatenated;
      if (i + 1 < current_level.size()) {
        concatenated = current_level[i] + current_level[i + 1];
      } else {

        concatenated = current_level[i] + current_level[i];
      }
      next_level.push_back(hash(concatenated));
    }
    current_level = next_level;
  }

  return current_level.front();
}

vector<string> get_blocks(const vector<string> &filenames) {
  vector<string> blocks;
  for (const auto &filename : filenames) {
    std::ifstream file(filename);
    auto content = string((std::istreambuf_iterator<char>(file)),
                          std::istreambuf_iterator<char>());
    blocks.push_back(hash(content));
  }

  return blocks;
}

string hash(const string &content) {
  EVP_MD_CTX *ctx = EVP_MD_CTX_new();
  const EVP_MD *md = EVP_sha256();
  unsigned char hash[EVP_MAX_MD_SIZE];
  unsigned int length;

  EVP_DigestInit_ex(ctx, md, nullptr);
  EVP_DigestUpdate(ctx, content.c_str(), content.size());
  EVP_DigestFinal_ex(ctx, hash, &length);
  EVP_MD_CTX_free(ctx);

  stringstream ss;
  for (unsigned int i = 0; i < length; ++i)
    ss << std::hex << std::setw(2) << std::setfill('0') << (int)hash[i];

  return ss.str();
}

vector<string> get_filenames(const path &directory) {
  vector<string> filenames;
  for (const auto &entry : recursive_directory_iterator(directory)) {
    if (entry.is_directory() or entry.path().filename() == "CHECKSUM")
      continue;

    filenames.push_back(entry.path().string());
  }

  return filenames;
}
} // namespace merkletree

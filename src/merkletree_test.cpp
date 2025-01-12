#include <catch2/catch_session.hpp>
#include <catch2/catch_test_macros.hpp>
#include <filesystem>
#include <string>
#include <vector>

#include "merkletree.h"

using std::string;
using std::vector;
using std::filesystem::current_path;
using std::filesystem::path;

const path articles_dir =
    current_path().parent_path() / "assets/fixtures/articles";

TEST_CASE("hash function") {
  REQUIRE(merkletree::hash("1") ==
          "6b86b273ff34fce19d6b804eff5a3f5747ada4eaa22f1d49c01e52ddb7875b4b");
  REQUIRE(merkletree::hash("2") ==
          "d4735e3a265e16eee03f59718b9b5d03019c07d8b6c51f90da3a666eec13ab35");
  REQUIRE(merkletree::hash("3") ==
          "4e07408562bedb8b60ce05c1decfe3ad16b72230967de01f640b7e4729b49fce");
}

TEST_CASE("Test get_filenames") {
  auto filenames = merkletree::get_filenames(articles_dir);
  auto expected = 24;
  auto result = filenames.size();
  REQUIRE(result == expected);
}

TEST_CASE("Basic functionality") {
  string expected =
      "c19ce1b23fc9057eb072011d793ce33a47bb6fc3fe4cf9bf5d8f737abd3be0cb";
  vector<string> blocks = {"1", "2", "3", "4", "5"};
  for (auto &block : blocks)
    block = merkletree::hash(block);

  string result = merkletree::find_root(blocks);

  REQUIRE(expected == result);
}

TEST_CASE("Single data block") {
  vector<string> filenames = {(articles_dir / "day_1/post_1")};
  vector<string> blocks = merkletree::get_blocks(filenames);
  auto result = merkletree::find_root(blocks);

  auto expected =
      "df14eda6e74f15dbaad5974e15d90c9bdca9ba601527a93b8d8b6d472e868d16";

  REQUIRE(expected == result);
}

TEST_CASE("Odd number of data blocks") {
  string expected =
      "f3f1917304e3af565b827d1baa9fac18d5b287ae97adda22dc51a0aef900b787";
  vector<string> blocks = {"1", "2", "3"};

  for (auto &block : blocks)
    block = merkletree::hash(block);

  string result = merkletree::find_root(blocks);

  REQUIRE(expected == result);
}

int main(int argc, char *argv[]) {
  return Catch::Session().run(argc, argv);
}

class Deskclean < Formula
  desc "Keep your ~/Desktop tidy!"
  homepage "https://github.com/audy/deskclean"
  url "https://github.com/audy/deskclean/archive/refs/tags/0.1.1.tar.gz"
  version "0.0.1"
  sha256 ""
  license "MIT"

  depends_on "go" => :build

  def install
    system "go", "build", "deskclean.go"
    system "mkdir", bin
    system "mv", "deskclean", bin
  end

  test do
    system bin / "deskclean", "-h"
  end
end

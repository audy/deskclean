class Deskclean < Formula
  desc "Keep your ~/Desktop tidy!"
  homepage "https://github.com/audy/deskclean"
  url "https://github.com/audy/deskclean/archive/refs/tags/0.1.2.tar.gz"
  version "0.1.2"
  sha256 "3249b6553492ba3097f1a24da3ba936b81596bd2"
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

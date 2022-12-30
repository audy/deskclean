class Deskclean < Formula
  desc "Keep your ~/Desktop tidy!"
  homepage "https://github.com/audy/deskclean"
  url "https://github.com/audy/deskclean/archive/refs/tags/0.1.2.tar.gz"
  version "0.1.2"
  sha256 "c36ab3e307ee6cec01e913a3d0aedfb7742049368025ff3c8b284a65efa9fcd5"
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

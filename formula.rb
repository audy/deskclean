class Deskclean < Formula
  homepage "https://github.com/audy/deskclean"
  url "https://github.com/audy/deskclean/archive/0.1.0.tar.gz"

  depends_on "go" => :build

  def install
    system "gobuild.sh"
    bin.install ".gobuild/bin/deskclean" => "deskclean"
  end

  test do
    system "#{bin}/deskclean", "--help"
  end
end

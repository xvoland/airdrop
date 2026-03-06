class Airdrop < Formula
  desc "CLI Utility for Apple AirDrop"
  homepage "https://github.com/xvoland/airdrop"
  license "MIT"
  version "0.3.4"
  url "https://github.com/xvoland/airdrop/releases/download/v0.3.4/airdrop_darwin_arm64.tar.gz"
  sha256 "0e85e9908186936f11d7bc104c9f2d9200d9561cabd44869cdfb66facc9886e4"

  def install
    bin.install "airdrop"
  end

  test do
    system "#{bin}/airdrop", "--version"
  end
end

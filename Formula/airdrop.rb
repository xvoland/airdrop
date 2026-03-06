class Airdrop < Formula
  desc "CLI Utility for Apple AirDrop"
  homepage "https://github.com/xvoland/airdrop"
  license "MIT"
  version "0.3.3"
  url "https://github.com/xvoland/airdrop/releases/download/v0.3.3/airdrop_darwin_universal.tar.gz"
  sha256 "e7a8990276b73f07d39dafa827fc95adbbce265ea253ad4c085e7ac643d1fab0"

  def install
    bin.install "airdrop"
  end

  test do
    system "#{bin}/airdrop", "--version"
  end
end

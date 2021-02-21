# typed: false
# frozen_string_literal: true

# This file was generated by GoReleaser. DO NOT EDIT.
class Serve < Formula
  desc "Simple HTTP file server."
  homepage "https://4d63.com/serve"
  version "1.1.0"
  bottle :unneeded

  if OS.mac?
    url "https://github.com/leighmcculloch/serve/releases/download/v1.1.0/serve_1.1.0_darwin_amd64.tar.gz"
    sha256 "e15afac5f3677b0103fe9129929bec91b5157e1248551ab3223bb036221f81b5"
  end
  if OS.linux? && Hardware::CPU.intel?
    url "https://github.com/leighmcculloch/serve/releases/download/v1.1.0/serve_1.1.0_linux_amd64.tar.gz"
    sha256 "1c27aa591f86cd52126244f2f5cb409a3518fc5f7b3d12191ac02ac5622fe856"
  end

  def install
    bin.install "serve"
  end
end

# typed: false
# frozen_string_literal: true

class Calendarn < Formula
  desc "CLI tool to generate Nepali and English calendars with date conversion"
  homepage "https://github.com/samit22/calendarN"
  url "https://github.com/samit22/calendarN/archive/refs/tags/v1.3.3.tar.gz"
  sha256 "REPLACE_WITH_SHA256"
  license "MIT"
  head "https://github.com/samit22/calendarN.git", branch: "main"

  depends_on "go" => :build

  def install
    # Read version from .version file
    app_version = File.read(".version").strip

    ldflags = %W[
      -s -w
      -X main.version=#{app_version}
    ]
    system "go", "build", *std_go_args(ldflags:), "-o", bin/"calendarN"
  end

  test do
    assert_match "calendarN", shell_output("#{bin}/calendarN --help")
    assert_match version.to_s, shell_output("#{bin}/calendarN version")
  end
end


class Nex < Formula
  desc "Simple server for serving files and directory listings over TCP"
  homepage "https://github.com/ShreyanJain9/nex-server"
  url "https://github.com/ShreyanJain9/nex-server/archive/refs/heads/main.tar.gz"
  sha256 "efc2dbf955abfd63df4df471c6f6ad778c4498d103d4b5bb3ba3ab08e1e1c5fb"
  license "AGPLv3"

  depends_on "go" => :build

  def install
    system "go", "build", "-o", "nex", "main.go"
    bin.install "nex"
  end

  def caveats
    <<~EOS
      The Nex Server can be configured using the configuration file at:
        #{etc}/nex/config.toml
    EOS
  end

  plist_options manual: "nex"

  def plist
    '
    <?xml version="1.0" encoding="UTF-8"?>
    <plist version="1.0">
      <dict>
        <key>Label</key>
        <string>#{plist_name}</string>
        <key>ProgramArguments</key>
        <array>
          <string>#{opt_bin}/nex</string>
        </array>
        <key>RunAtLoad</key>
        <true/>
      </dict>
    </plist>
  '
  end
end

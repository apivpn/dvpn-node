package v2ray

import (
	"strings"
)

var (
	configTemplate = strings.TrimSpace(`
{
    "api": {
        "services": [
            "HandlerService",
            "StatsService"
        ],
        "tag": "api"
    },
    "inbounds": [
        {
            "listen": "127.0.0.1",
            "port": 23,
            "protocol": "dokodemo-door",
            "settings": {
                "address": "127.0.0.1"
            },
            "tag": "api"
        },
        {
            "port": "{{ .VMess.ListenPort }}",
            "protocol": "vmess",
            "streamSettings": {
                "network": "{{ .VMess.Transport }}"
            },
            "tag": "{{ .VMess.Transport }}"
        },
        {
			"tag": "ws",
            "port": "443",
            "protocol": "vmess",
            "streamSettings": {
                "network": "ws",
                "security": "tls",
                "wsSettings": {
                    "path": "/ws"
                },
                "tlsSettings": {
                    "allowInsecure": true,
                    "certificates": [
                        {
                            "certificateFile": "{{ .VMess.Cert }}",
                            "keyFile": "{{ .VMess.Key }}"
                        }
                    ]
                }
            }
        }
    ],
    "log": {
        "loglevel": "debug"
    },
    "outbounds": [
        {
            "protocol": "freedom"
        }
    ],
    "policy": {
        "levels": {
            "0": {
                "statsUserDownlink": true,
                "statsUserUplink": true
            }
        }
    },
    "routing": {
        "rules": [
            {
                "inboundTag": [
                    "api"
                ],
                "outboundTag": "api",
                "type": "field"
            }
        ]
    },
    "stats": {},
    "transport": {
        "dsSettings": {},
        "grpcSettings": {},
        "gunSettings": {},
        "httpSettings": {},
        "kcpSettings": {},
        "quicSettings": {
            "security": "chacha20-poly1305"
        },
        "tcpSettings": {},
        "wsSettings": {}
    }
}
	`)
)

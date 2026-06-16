#!/bin/sh

set -e

# --- 环境变量检查 (保持不变) ---
: "${DB_HOST:?错误：必须设置 DB_HOST 环境变量}"
: "${DB_PORT:?错误：必须设置 DB_PORT 环境变量}"
: "${DB_USER:?错误：必须设置 DB_USER 环境变量}"
: "${DB_PASS:?错误：必须设置 DB_PASS 环境变量}"
: "${DB_NAME:?错误：必须设置 DB_NAME 环境变量}"
: "${BIND_ADDR:?错误：必须设置 BIND_ADDR 环境变量}"

echo "启动 komari 服务..."
echo "监听地址: ${BIND_ADDR}"
echo "数据库主机: ${DB_HOST}"

# --- 【核心修正】---
# 直接使用环境变量来构建命令，而不是依赖 "$@"
exec ./komari server \
    --db-type "${DB_TYPE}" \
    --db-host "${DB_HOST}" \
    --db-port "${DB_PORT}" \
    --db-user "${DB_USER}" \
    --db-pass "${DB_PASS}" \
    --db-name "${DB_NAME}" \
    -l "${BIND_ADDR}"

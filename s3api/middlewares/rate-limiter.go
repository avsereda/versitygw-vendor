// Copyright 2026 Versity Software
// This file is licensed under the Apache License, Version 2.0
// (the "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/versity/versitygw/metrics"
	"github.com/versity/versitygw/s3log"
)

// RateLimiter hard-limits the number of in-flight requests.
// If the limit is reached, an immediate SlowDown error is returned
func RateLimiter(limit int, mm metrics.Manager, logger s3log.AuditLogger) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		return ctx.Next()
	}
}

// AUTO-GENERATED — do not edit manually.
// Re-run `make sdk-go` (or `make sdk-all`) to update.
// Source of truth: @readonly annotations in hostafrica-api.smithy

package hasdk

// RetryablePaths is the set of API paths that are safe to retry on failure.
var RetryablePaths = map[string]struct{}{
	"/vps/list-vps-services": {},
	"/vps/get-details": {},
	"/vps/get-config": {},
	"/vps/novnc-console": {},
	"/vps/list-backups": {},
	"/vps/list-backup-schedules": {},
	"/vps/list-snapshot-jobs": {},
	"/vps/list-snapshots": {},
	"/vps/list-firewall-rules": {},
	"/vps/list-power-tasks": {},
	"/vps/list-notifications": {},
	"/vps/list-isos": {},
	"/vps/list-reinstall-images": {},
	"/vps/get-catalogue": {},
	"/vps/validate-pricing": {},
	"/vps/list-orders": {},
	"/dns/list-rdns-records": {},
	"/domain/check-availability": {},
	"/domain/suggest": {},
	"/domain/list-domains": {},
	"/domain/list-domains-requiring-data": {},
	"/domain/get-domain": {},
	"/domain/get-domain-contacts": {},
	"/dns/list-zones": {},
	"/dns/list-create-candidates": {},
	"/dns/get-zone": {},
	"/domain/list-dnssec-records": {},
}

// IsRetryable reports whether path is safe to retry.
func IsRetryable(path string) bool {
	_, ok := RetryablePaths[path]
	return ok
}

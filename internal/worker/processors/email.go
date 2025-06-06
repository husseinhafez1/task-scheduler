package processors

import (
	"context"
	"encoding/json"
	"log"
	"task/internal/metrics"

	"github.com/prometheus/client_golang/prometheus"
)

type EmailJobData struct {
	To      string `json:"to"`
	Subject string `json:"subject"`
	Body    string `json:"body"`
}

func ProcessEmailJob(ctx context.Context, data json.RawMessage) bool {
	// Start timing the email job
	timer := prometheus.NewTimer(metrics.JobDurationByType.WithLabelValues("email"))
	defer timer.ObserveDuration()

	var emailData EmailJobData
	if err := json.Unmarshal(data, &emailData); err != nil {
		log.Printf("Error parsing email job data: %v", err)
		metrics.JobsFailedByType.WithLabelValues("email").Inc()
		return false
	}

	// Validate email data
	if emailData.To == "" || emailData.Subject == "" || emailData.Body == "" {
		log.Printf("Invalid email data: missing required fields")
		metrics.JobsFailedByType.WithLabelValues("email").Inc()
		return false
	}

	// TODO: Replace with actual email sending logic
	// For now, we'll just simulate sending an email
	log.Printf("Sending email to %s: %s", emailData.To, emailData.Subject)

	// Simulate some processing time
	select {
	case <-ctx.Done():
		log.Printf("Email job cancelled due to timeout")
		metrics.JobsFailedByType.WithLabelValues("email").Inc()
		return false
	default:
		// Continue processing
	}

	// Simulate successful email sending
	log.Printf("Email sent successfully to %s", emailData.To)
	return true
}

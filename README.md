# Webhook Ingestion Service

A production-style Go backend service for securely ingesting and persisting webhook events with strong correctness guarantees.

This project focuses on **reliability, security, and system correctness**, not UI or product features.

---

## Overview

This service implements a real-world webhook ingestion pattern similar to systems used by Stripe, GitHub, and other SaaS platforms.

It accepts incoming webhook events, verifies their authenticity using HMAC signatures, and persists them durably for later asynchronous processing.
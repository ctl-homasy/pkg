// pkg/middleware/security.go

package middleware

import (
	"context"
	"fmt"
	"log"
	"runtime/debug"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// LoggingInterceptor creates a gRPC interceptor for request logging
func LoggingInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		start := time.Now()
		
		// Extract request metadata
		md, _ := metadata.FromIncomingContext(ctx)
		
		// Log request
		log.Printf("Request: method=%s, metadata=%v", info.FullMethod, md)
		
		// Handle the request
		resp, err := handler(ctx, req)
		
		// Log response
		duration := time.Since(start)
		if err != nil {
			log.Printf("Response: method=%s, error=%v, duration=%s", info.FullMethod, err, duration)
		} else {
			log.Printf("Response: method=%s, duration=%s", info.FullMethod, duration)
		}
		
		return resp, err
	}
}

// RecoveryInterceptor creates a gRPC interceptor for panic recovery
func RecoveryInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("Panic recovered: %v\n%s", r, debug.Stack())
				err = status.Errorf(codes.Internal, "Internal server error")
			}
		}()
		
		return handler(ctx, req)
	}
}

// IPFilterInterceptor creates a gRPC interceptor for IP filtering
func IPFilterInterceptor(allowedIPs []string) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		// Extract client IP from context or metadata
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, status.Errorf(codes.Unauthenticated, "missing metadata")
		}
		
		// Check if x-forwarded-for or x-real-ip header exists
		var clientIP string
		if ips := md.Get("x-forwarded-for"); len(ips) > 0 {
			clientIP = ips[0]
		} else if ips := md.Get("x-real-ip"); len(ips) > 0 {
			clientIP = ips[0]
		}
		
		// If no allowed IPs are specified, allow all
		if len(allowedIPs) == 0 {
			return handler(ctx, req)
		}
		
		// Check if client IP is in the allowed list
		allowed := false
		for _, ip := range allowedIPs {
			if ip == clientIP {
				allowed = true
				break
			}
		}
		
		if !allowed {
			return nil, status.Errorf(codes.PermissionDenied, "access denied from IP: %s", clientIP)
		}
		
		return handler(ctx, req)
	}
}

// RequestIDInterceptor adds a unique request ID to the context
func RequestIDInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		// Generate a unique request ID
		requestID := fmt.Sprintf("%d", time.Now().UnixNano())
		
		// Add request ID to outgoing metadata
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			md = metadata.New(nil)
		}
		md = md.Copy()
		md.Set("x-request-id", requestID)
		
		// Create a new context with the request ID
		newCtx := metadata.NewIncomingContext(ctx, md)
		
		// Log the request ID
		log.Printf("Request ID: %s, Method: %s", requestID, info.FullMethod)
		
		return handler(newCtx, req)
	}
}
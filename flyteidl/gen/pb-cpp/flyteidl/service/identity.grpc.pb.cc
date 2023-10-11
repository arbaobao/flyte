// Generated by the gRPC C++ plugin.
// If you make any local change, they will be lost.
// source: flyteidl/service/identity.proto

#include "flyteidl/service/identity.pb.h"
#include "flyteidl/service/identity.grpc.pb.h"

#include <functional>
#include <grpcpp/impl/codegen/async_stream.h>
#include <grpcpp/impl/codegen/async_unary_call.h>
#include <grpcpp/impl/codegen/channel_interface.h>
#include <grpcpp/impl/codegen/client_unary_call.h>
#include <grpcpp/impl/codegen/client_callback.h>
#include <grpcpp/impl/codegen/method_handler_impl.h>
#include <grpcpp/impl/codegen/rpc_service_method.h>
#include <grpcpp/impl/codegen/server_callback.h>
#include <grpcpp/impl/codegen/service_type.h>
#include <grpcpp/impl/codegen/sync_stream.h>
namespace flyteidl {
namespace service {

static const char* IdentityService_method_names[] = {
  "/flyteidl.service.IdentityService/UserInfo",
};

std::unique_ptr< IdentityService::Stub> IdentityService::NewStub(const std::shared_ptr< ::grpc::ChannelInterface>& channel, const ::grpc::StubOptions& options) {
  (void)options;
  std::unique_ptr< IdentityService::Stub> stub(new IdentityService::Stub(channel));
  return stub;
}

IdentityService::Stub::Stub(const std::shared_ptr< ::grpc::ChannelInterface>& channel)
  : channel_(channel), rpcmethod_UserInfo_(IdentityService_method_names[0], ::grpc::internal::RpcMethod::NORMAL_RPC, channel)
  {}

::grpc::Status IdentityService::Stub::UserInfo(::grpc::ClientContext* context, const ::flyteidl::service::UserInfoRequest& request, ::flyteidl::service::UserInfoResponse* response) {
  return ::grpc::internal::BlockingUnaryCall(channel_.get(), rpcmethod_UserInfo_, context, request, response);
}

void IdentityService::Stub::experimental_async::UserInfo(::grpc::ClientContext* context, const ::flyteidl::service::UserInfoRequest* request, ::flyteidl::service::UserInfoResponse* response, std::function<void(::grpc::Status)> f) {
  ::grpc::internal::CallbackUnaryCall(stub_->channel_.get(), stub_->rpcmethod_UserInfo_, context, request, response, std::move(f));
}

void IdentityService::Stub::experimental_async::UserInfo(::grpc::ClientContext* context, const ::grpc::ByteBuffer* request, ::flyteidl::service::UserInfoResponse* response, std::function<void(::grpc::Status)> f) {
  ::grpc::internal::CallbackUnaryCall(stub_->channel_.get(), stub_->rpcmethod_UserInfo_, context, request, response, std::move(f));
}

void IdentityService::Stub::experimental_async::UserInfo(::grpc::ClientContext* context, const ::flyteidl::service::UserInfoRequest* request, ::flyteidl::service::UserInfoResponse* response, ::grpc::experimental::ClientUnaryReactor* reactor) {
  ::grpc::internal::ClientCallbackUnaryFactory::Create(stub_->channel_.get(), stub_->rpcmethod_UserInfo_, context, request, response, reactor);
}

void IdentityService::Stub::experimental_async::UserInfo(::grpc::ClientContext* context, const ::grpc::ByteBuffer* request, ::flyteidl::service::UserInfoResponse* response, ::grpc::experimental::ClientUnaryReactor* reactor) {
  ::grpc::internal::ClientCallbackUnaryFactory::Create(stub_->channel_.get(), stub_->rpcmethod_UserInfo_, context, request, response, reactor);
}

::grpc::ClientAsyncResponseReader< ::flyteidl::service::UserInfoResponse>* IdentityService::Stub::AsyncUserInfoRaw(::grpc::ClientContext* context, const ::flyteidl::service::UserInfoRequest& request, ::grpc::CompletionQueue* cq) {
  return ::grpc::internal::ClientAsyncResponseReaderFactory< ::flyteidl::service::UserInfoResponse>::Create(channel_.get(), cq, rpcmethod_UserInfo_, context, request, true);
}

::grpc::ClientAsyncResponseReader< ::flyteidl::service::UserInfoResponse>* IdentityService::Stub::PrepareAsyncUserInfoRaw(::grpc::ClientContext* context, const ::flyteidl::service::UserInfoRequest& request, ::grpc::CompletionQueue* cq) {
  return ::grpc::internal::ClientAsyncResponseReaderFactory< ::flyteidl::service::UserInfoResponse>::Create(channel_.get(), cq, rpcmethod_UserInfo_, context, request, false);
}

IdentityService::Service::Service() {
  AddMethod(new ::grpc::internal::RpcServiceMethod(
      IdentityService_method_names[0],
      ::grpc::internal::RpcMethod::NORMAL_RPC,
      new ::grpc::internal::RpcMethodHandler< IdentityService::Service, ::flyteidl::service::UserInfoRequest, ::flyteidl::service::UserInfoResponse>(
          std::mem_fn(&IdentityService::Service::UserInfo), this)));
}

IdentityService::Service::~Service() {
}

::grpc::Status IdentityService::Service::UserInfo(::grpc::ServerContext* context, const ::flyteidl::service::UserInfoRequest* request, ::flyteidl::service::UserInfoResponse* response) {
  (void) context;
  (void) request;
  (void) response;
  return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
}


}  // namespace flyteidl
}  // namespace service

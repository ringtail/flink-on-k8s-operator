// Copyright 2018 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// AUTO-GENERATED CODE. DO NOT EDIT.

package vision

import (
	emptypb "github.com/golang/protobuf/ptypes/empty"
	visionpb "google.golang.org/genproto/googleapis/cloud/vision/v1"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
)

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
	"testing"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"google.golang.org/api/option"
	status "google.golang.org/genproto/googleapis/rpc/status"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	gstatus "google.golang.org/grpc/status"
)

var _ = io.EOF
var _ = ptypes.MarshalAny
var _ status.Status

type mockProductSearchServer struct {
	// Embed for forward compatibility.
	// Tests will keep working if more methods are added
	// in the future.
	visionpb.ProductSearchServer

	reqs []proto.Message

	// If set, all calls return this error.
	err error

	// responses to return if err == nil
	resps []proto.Message
}

func (s *mockProductSearchServer) CreateProductSet(ctx context.Context, req *visionpb.CreateProductSetRequest) (*visionpb.ProductSet, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	if xg := md["x-goog-api-client"]; len(xg) == 0 || !strings.Contains(xg[0], "gl-go/") {
		return nil, fmt.Errorf("x-goog-api-client = %v, expected gl-go key", xg)
	}
	s.reqs = append(s.reqs, req)
	if s.err != nil {
		return nil, s.err
	}
	return s.resps[0].(*visionpb.ProductSet), nil
}

func (s *mockProductSearchServer) ListProductSets(ctx context.Context, req *visionpb.ListProductSetsRequest) (*visionpb.ListProductSetsResponse, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	if xg := md["x-goog-api-client"]; len(xg) == 0 || !strings.Contains(xg[0], "gl-go/") {
		return nil, fmt.Errorf("x-goog-api-client = %v, expected gl-go key", xg)
	}
	s.reqs = append(s.reqs, req)
	if s.err != nil {
		return nil, s.err
	}
	return s.resps[0].(*visionpb.ListProductSetsResponse), nil
}

func (s *mockProductSearchServer) GetProductSet(ctx context.Context, req *visionpb.GetProductSetRequest) (*visionpb.ProductSet, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	if xg := md["x-goog-api-client"]; len(xg) == 0 || !strings.Contains(xg[0], "gl-go/") {
		return nil, fmt.Errorf("x-goog-api-client = %v, expected gl-go key", xg)
	}
	s.reqs = append(s.reqs, req)
	if s.err != nil {
		return nil, s.err
	}
	return s.resps[0].(*visionpb.ProductSet), nil
}

func (s *mockProductSearchServer) UpdateProductSet(ctx context.Context, req *visionpb.UpdateProductSetRequest) (*visionpb.ProductSet, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	if xg := md["x-goog-api-client"]; len(xg) == 0 || !strings.Contains(xg[0], "gl-go/") {
		return nil, fmt.Errorf("x-goog-api-client = %v, expected gl-go key", xg)
	}
	s.reqs = append(s.reqs, req)
	if s.err != nil {
		return nil, s.err
	}
	return s.resps[0].(*visionpb.ProductSet), nil
}

func (s *mockProductSearchServer) DeleteProductSet(ctx context.Context, req *visionpb.DeleteProductSetRequest) (*emptypb.Empty, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	if xg := md["x-goog-api-client"]; len(xg) == 0 || !strings.Contains(xg[0], "gl-go/") {
		return nil, fmt.Errorf("x-goog-api-client = %v, expected gl-go key", xg)
	}
	s.reqs = append(s.reqs, req)
	if s.err != nil {
		return nil, s.err
	}
	return s.resps[0].(*emptypb.Empty), nil
}

func (s *mockProductSearchServer) CreateProduct(ctx context.Context, req *visionpb.CreateProductRequest) (*visionpb.Product, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	if xg := md["x-goog-api-client"]; len(xg) == 0 || !strings.Contains(xg[0], "gl-go/") {
		return nil, fmt.Errorf("x-goog-api-client = %v, expected gl-go key", xg)
	}
	s.reqs = append(s.reqs, req)
	if s.err != nil {
		return nil, s.err
	}
	return s.resps[0].(*visionpb.Product), nil
}

func (s *mockProductSearchServer) ListProducts(ctx context.Context, req *visionpb.ListProductsRequest) (*visionpb.ListProductsResponse, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	if xg := md["x-goog-api-client"]; len(xg) == 0 || !strings.Contains(xg[0], "gl-go/") {
		return nil, fmt.Errorf("x-goog-api-client = %v, expected gl-go key", xg)
	}
	s.reqs = append(s.reqs, req)
	if s.err != nil {
		return nil, s.err
	}
	return s.resps[0].(*visionpb.ListProductsResponse), nil
}

func (s *mockProductSearchServer) GetProduct(ctx context.Context, req *visionpb.GetProductRequest) (*visionpb.Product, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	if xg := md["x-goog-api-client"]; len(xg) == 0 || !strings.Contains(xg[0], "gl-go/") {
		return nil, fmt.Errorf("x-goog-api-client = %v, expected gl-go key", xg)
	}
	s.reqs = append(s.reqs, req)
	if s.err != nil {
		return nil, s.err
	}
	return s.resps[0].(*visionpb.Product), nil
}

func (s *mockProductSearchServer) UpdateProduct(ctx context.Context, req *visionpb.UpdateProductRequest) (*visionpb.Product, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	if xg := md["x-goog-api-client"]; len(xg) == 0 || !strings.Contains(xg[0], "gl-go/") {
		return nil, fmt.Errorf("x-goog-api-client = %v, expected gl-go key", xg)
	}
	s.reqs = append(s.reqs, req)
	if s.err != nil {
		return nil, s.err
	}
	return s.resps[0].(*visionpb.Product), nil
}

func (s *mockProductSearchServer) DeleteProduct(ctx context.Context, req *visionpb.DeleteProductRequest) (*emptypb.Empty, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	if xg := md["x-goog-api-client"]; len(xg) == 0 || !strings.Contains(xg[0], "gl-go/") {
		return nil, fmt.Errorf("x-goog-api-client = %v, expected gl-go key", xg)
	}
	s.reqs = append(s.reqs, req)
	if s.err != nil {
		return nil, s.err
	}
	return s.resps[0].(*emptypb.Empty), nil
}

func (s *mockProductSearchServer) CreateReferenceImage(ctx context.Context, req *visionpb.CreateReferenceImageRequest) (*visionpb.ReferenceImage, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	if xg := md["x-goog-api-client"]; len(xg) == 0 || !strings.Contains(xg[0], "gl-go/") {
		return nil, fmt.Errorf("x-goog-api-client = %v, expected gl-go key", xg)
	}
	s.reqs = append(s.reqs, req)
	if s.err != nil {
		return nil, s.err
	}
	return s.resps[0].(*visionpb.ReferenceImage), nil
}

func (s *mockProductSearchServer) DeleteReferenceImage(ctx context.Context, req *visionpb.DeleteReferenceImageRequest) (*emptypb.Empty, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	if xg := md["x-goog-api-client"]; len(xg) == 0 || !strings.Contains(xg[0], "gl-go/") {
		return nil, fmt.Errorf("x-goog-api-client = %v, expected gl-go key", xg)
	}
	s.reqs = append(s.reqs, req)
	if s.err != nil {
		return nil, s.err
	}
	return s.resps[0].(*emptypb.Empty), nil
}

func (s *mockProductSearchServer) ListReferenceImages(ctx context.Context, req *visionpb.ListReferenceImagesRequest) (*visionpb.ListReferenceImagesResponse, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	if xg := md["x-goog-api-client"]; len(xg) == 0 || !strings.Contains(xg[0], "gl-go/") {
		return nil, fmt.Errorf("x-goog-api-client = %v, expected gl-go key", xg)
	}
	s.reqs = append(s.reqs, req)
	if s.err != nil {
		return nil, s.err
	}
	return s.resps[0].(*visionpb.ListReferenceImagesResponse), nil
}

func (s *mockProductSearchServer) GetReferenceImage(ctx context.Context, req *visionpb.GetReferenceImageRequest) (*visionpb.ReferenceImage, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	if xg := md["x-goog-api-client"]; len(xg) == 0 || !strings.Contains(xg[0], "gl-go/") {
		return nil, fmt.Errorf("x-goog-api-client = %v, expected gl-go key", xg)
	}
	s.reqs = append(s.reqs, req)
	if s.err != nil {
		return nil, s.err
	}
	return s.resps[0].(*visionpb.ReferenceImage), nil
}

func (s *mockProductSearchServer) AddProductToProductSet(ctx context.Context, req *visionpb.AddProductToProductSetRequest) (*emptypb.Empty, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	if xg := md["x-goog-api-client"]; len(xg) == 0 || !strings.Contains(xg[0], "gl-go/") {
		return nil, fmt.Errorf("x-goog-api-client = %v, expected gl-go key", xg)
	}
	s.reqs = append(s.reqs, req)
	if s.err != nil {
		return nil, s.err
	}
	return s.resps[0].(*emptypb.Empty), nil
}

func (s *mockProductSearchServer) RemoveProductFromProductSet(ctx context.Context, req *visionpb.RemoveProductFromProductSetRequest) (*emptypb.Empty, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	if xg := md["x-goog-api-client"]; len(xg) == 0 || !strings.Contains(xg[0], "gl-go/") {
		return nil, fmt.Errorf("x-goog-api-client = %v, expected gl-go key", xg)
	}
	s.reqs = append(s.reqs, req)
	if s.err != nil {
		return nil, s.err
	}
	return s.resps[0].(*emptypb.Empty), nil
}

func (s *mockProductSearchServer) ListProductsInProductSet(ctx context.Context, req *visionpb.ListProductsInProductSetRequest) (*visionpb.ListProductsInProductSetResponse, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	if xg := md["x-goog-api-client"]; len(xg) == 0 || !strings.Contains(xg[0], "gl-go/") {
		return nil, fmt.Errorf("x-goog-api-client = %v, expected gl-go key", xg)
	}
	s.reqs = append(s.reqs, req)
	if s.err != nil {
		return nil, s.err
	}
	return s.resps[0].(*visionpb.ListProductsInProductSetResponse), nil
}

func (s *mockProductSearchServer) ImportProductSets(ctx context.Context, req *visionpb.ImportProductSetsRequest) (*longrunningpb.Operation, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	if xg := md["x-goog-api-client"]; len(xg) == 0 || !strings.Contains(xg[0], "gl-go/") {
		return nil, fmt.Errorf("x-goog-api-client = %v, expected gl-go key", xg)
	}
	s.reqs = append(s.reqs, req)
	if s.err != nil {
		return nil, s.err
	}
	return s.resps[0].(*longrunningpb.Operation), nil
}

type mockImageAnnotatorServer struct {
	// Embed for forward compatibility.
	// Tests will keep working if more methods are added
	// in the future.
	visionpb.ImageAnnotatorServer

	reqs []proto.Message

	// If set, all calls return this error.
	err error

	// responses to return if err == nil
	resps []proto.Message
}

func (s *mockImageAnnotatorServer) BatchAnnotateImages(ctx context.Context, req *visionpb.BatchAnnotateImagesRequest) (*visionpb.BatchAnnotateImagesResponse, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	if xg := md["x-goog-api-client"]; len(xg) == 0 || !strings.Contains(xg[0], "gl-go/") {
		return nil, fmt.Errorf("x-goog-api-client = %v, expected gl-go key", xg)
	}
	s.reqs = append(s.reqs, req)
	if s.err != nil {
		return nil, s.err
	}
	return s.resps[0].(*visionpb.BatchAnnotateImagesResponse), nil
}

func (s *mockImageAnnotatorServer) AsyncBatchAnnotateFiles(ctx context.Context, req *visionpb.AsyncBatchAnnotateFilesRequest) (*longrunningpb.Operation, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	if xg := md["x-goog-api-client"]; len(xg) == 0 || !strings.Contains(xg[0], "gl-go/") {
		return nil, fmt.Errorf("x-goog-api-client = %v, expected gl-go key", xg)
	}
	s.reqs = append(s.reqs, req)
	if s.err != nil {
		return nil, s.err
	}
	return s.resps[0].(*longrunningpb.Operation), nil
}

// clientOpt is the option tests should use to connect to the test server.
// It is initialized by TestMain.
var clientOpt option.ClientOption

var (
	mockProductSearch  mockProductSearchServer
	mockImageAnnotator mockImageAnnotatorServer
)

func TestMain(m *testing.M) {
	flag.Parse()

	serv := grpc.NewServer()
	visionpb.RegisterProductSearchServer(serv, &mockProductSearch)
	visionpb.RegisterImageAnnotatorServer(serv, &mockImageAnnotator)

	lis, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		log.Fatal(err)
	}
	go serv.Serve(lis)

	conn, err := grpc.Dial(lis.Addr().String(), grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	clientOpt = option.WithGRPCConn(conn)

	os.Exit(m.Run())
}

func TestProductSearchCreateProduct(t *testing.T) {
	var name string = "name3373707"
	var displayName string = "displayName1615086568"
	var description string = "description-1724546052"
	var productCategory string = "productCategory-1607451058"
	var expectedResponse = &visionpb.Product{
		Name:            name,
		DisplayName:     displayName,
		Description:     description,
		ProductCategory: productCategory,
	}

	mockProductSearch.err = nil
	mockProductSearch.reqs = nil

	mockProductSearch.resps = append(mockProductSearch.resps[:0], expectedResponse)

	var formattedParent string = fmt.Sprintf("projects/%s/locations/%s", "[PROJECT]", "[LOCATION]")
	var product *visionpb.Product = &visionpb.Product{}
	var request = &visionpb.CreateProductRequest{
		Parent:  formattedParent,
		Product: product,
	}

	c, err := NewProductSearchClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := c.CreateProduct(context.Background(), request)

	if err != nil {
		t.Fatal(err)
	}

	if want, got := request, mockProductSearch.reqs[0]; !proto.Equal(want, got) {
		t.Errorf("wrong request %q, want %q", got, want)
	}

	if want, got := expectedResponse, resp; !proto.Equal(want, got) {
		t.Errorf("wrong response %q, want %q)", got, want)
	}
}

func TestProductSearchCreateProductError(t *testing.T) {
	errCode := codes.PermissionDenied
	mockProductSearch.err = gstatus.Error(errCode, "test error")

	var formattedParent string = fmt.Sprintf("projects/%s/locations/%s", "[PROJECT]", "[LOCATION]")
	var product *visionpb.Product = &visionpb.Product{}
	var request = &visionpb.CreateProductRequest{
		Parent:  formattedParent,
		Product: product,
	}

	c, err := NewProductSearchClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := c.CreateProduct(context.Background(), request)

	if st, ok := gstatus.FromError(err); !ok {
		t.Errorf("got error %v, expected grpc error", err)
	} else if c := st.Code(); c != errCode {
		t.Errorf("got error code %q, want %q", c, errCode)
	}
	_ = resp
}
func TestProductSearchListProducts(t *testing.T) {
	var nextPageToken string = ""
	var productsElement *visionpb.Product = &visionpb.Product{}
	var products = []*visionpb.Product{productsElement}
	var expectedResponse = &visionpb.ListProductsResponse{
		NextPageToken: nextPageToken,
		Products:      products,
	}

	mockProductSearch.err = nil
	mockProductSearch.reqs = nil

	mockProductSearch.resps = append(mockProductSearch.resps[:0], expectedResponse)

	var formattedParent string = fmt.Sprintf("projects/%s/locations/%s", "[PROJECT]", "[LOCATION]")
	var request = &visionpb.ListProductsRequest{
		Parent: formattedParent,
	}

	c, err := NewProductSearchClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := c.ListProducts(context.Background(), request).Next()

	if err != nil {
		t.Fatal(err)
	}

	if want, got := request, mockProductSearch.reqs[0]; !proto.Equal(want, got) {
		t.Errorf("wrong request %q, want %q", got, want)
	}

	want := (interface{})(expectedResponse.Products[0])
	got := (interface{})(resp)
	var ok bool

	switch want := (want).(type) {
	case proto.Message:
		ok = proto.Equal(want, got.(proto.Message))
	default:
		ok = want == got
	}
	if !ok {
		t.Errorf("wrong response %q, want %q)", got, want)
	}
}

func TestProductSearchListProductsError(t *testing.T) {
	errCode := codes.PermissionDenied
	mockProductSearch.err = gstatus.Error(errCode, "test error")

	var formattedParent string = fmt.Sprintf("projects/%s/locations/%s", "[PROJECT]", "[LOCATION]")
	var request = &visionpb.ListProductsRequest{
		Parent: formattedParent,
	}

	c, err := NewProductSearchClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := c.ListProducts(context.Background(), request).Next()

	if st, ok := gstatus.FromError(err); !ok {
		t.Errorf("got error %v, expected grpc error", err)
	} else if c := st.Code(); c != errCode {
		t.Errorf("got error code %q, want %q", c, errCode)
	}
	_ = resp
}
func TestProductSearchGetProduct(t *testing.T) {
	var name2 string = "name2-1052831874"
	var displayName string = "displayName1615086568"
	var description string = "description-1724546052"
	var productCategory string = "productCategory-1607451058"
	var expectedResponse = &visionpb.Product{
		Name:            name2,
		DisplayName:     displayName,
		Description:     description,
		ProductCategory: productCategory,
	}

	mockProductSearch.err = nil
	mockProductSearch.reqs = nil

	mockProductSearch.resps = append(mockProductSearch.resps[:0], expectedResponse)

	var formattedName string = fmt.Sprintf("projects/%s/locations/%s/products/%s", "[PROJECT]", "[LOCATION]", "[PRODUCT]")
	var request = &visionpb.GetProductRequest{
		Name: formattedName,
	}

	c, err := NewProductSearchClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := c.GetProduct(context.Background(), request)

	if err != nil {
		t.Fatal(err)
	}

	if want, got := request, mockProductSearch.reqs[0]; !proto.Equal(want, got) {
		t.Errorf("wrong request %q, want %q", got, want)
	}

	if want, got := expectedResponse, resp; !proto.Equal(want, got) {
		t.Errorf("wrong response %q, want %q)", got, want)
	}
}

func TestProductSearchGetProductError(t *testing.T) {
	errCode := codes.PermissionDenied
	mockProductSearch.err = gstatus.Error(errCode, "test error")

	var formattedName string = fmt.Sprintf("projects/%s/locations/%s/products/%s", "[PROJECT]", "[LOCATION]", "[PRODUCT]")
	var request = &visionpb.GetProductRequest{
		Name: formattedName,
	}

	c, err := NewProductSearchClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := c.GetProduct(context.Background(), request)

	if st, ok := gstatus.FromError(err); !ok {
		t.Errorf("got error %v, expected grpc error", err)
	} else if c := st.Code(); c != errCode {
		t.Errorf("got error code %q, want %q", c, errCode)
	}
	_ = resp
}
func TestProductSearchUpdateProduct(t *testing.T) {
	var name string = "name3373707"
	var displayName string = "displayName1615086568"
	var description string = "description-1724546052"
	var productCategory string = "productCategory-1607451058"
	var expectedResponse = &visionpb.Product{
		Name:            name,
		DisplayName:     displayName,
		Description:     description,
		ProductCategory: productCategory,
	}

	mockProductSearch.err = nil
	mockProductSearch.reqs = nil

	mockProductSearch.resps = append(mockProductSearch.resps[:0], expectedResponse)

	var product *visionpb.Product = &visionpb.Product{}
	var request = &visionpb.UpdateProductRequest{
		Product: product,
	}

	c, err := NewProductSearchClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := c.UpdateProduct(context.Background(), request)

	if err != nil {
		t.Fatal(err)
	}

	if want, got := request, mockProductSearch.reqs[0]; !proto.Equal(want, got) {
		t.Errorf("wrong request %q, want %q", got, want)
	}

	if want, got := expectedResponse, resp; !proto.Equal(want, got) {
		t.Errorf("wrong response %q, want %q)", got, want)
	}
}

func TestProductSearchUpdateProductError(t *testing.T) {
	errCode := codes.PermissionDenied
	mockProductSearch.err = gstatus.Error(errCode, "test error")

	var product *visionpb.Product = &visionpb.Product{}
	var request = &visionpb.UpdateProductRequest{
		Product: product,
	}

	c, err := NewProductSearchClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := c.UpdateProduct(context.Background(), request)

	if st, ok := gstatus.FromError(err); !ok {
		t.Errorf("got error %v, expected grpc error", err)
	} else if c := st.Code(); c != errCode {
		t.Errorf("got error code %q, want %q", c, errCode)
	}
	_ = resp
}
func TestProductSearchDeleteProduct(t *testing.T) {
	var expectedResponse *emptypb.Empty = &emptypb.Empty{}

	mockProductSearch.err = nil
	mockProductSearch.reqs = nil

	mockProductSearch.resps = append(mockProductSearch.resps[:0], expectedResponse)

	var formattedName string = fmt.Sprintf("projects/%s/locations/%s/products/%s", "[PROJECT]", "[LOCATION]", "[PRODUCT]")
	var request = &visionpb.DeleteProductRequest{
		Name: formattedName,
	}

	c, err := NewProductSearchClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	err = c.DeleteProduct(context.Background(), request)

	if err != nil {
		t.Fatal(err)
	}

	if want, got := request, mockProductSearch.reqs[0]; !proto.Equal(want, got) {
		t.Errorf("wrong request %q, want %q", got, want)
	}

}

func TestProductSearchDeleteProductError(t *testing.T) {
	errCode := codes.PermissionDenied
	mockProductSearch.err = gstatus.Error(errCode, "test error")

	var formattedName string = fmt.Sprintf("projects/%s/locations/%s/products/%s", "[PROJECT]", "[LOCATION]", "[PRODUCT]")
	var request = &visionpb.DeleteProductRequest{
		Name: formattedName,
	}

	c, err := NewProductSearchClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	err = c.DeleteProduct(context.Background(), request)

	if st, ok := gstatus.FromError(err); !ok {
		t.Errorf("got error %v, expected grpc error", err)
	} else if c := st.Code(); c != errCode {
		t.Errorf("got error code %q, want %q", c, errCode)
	}
}
func TestProductSearchListReferenceImages(t *testing.T) {
	var pageSize int32 = 883849137
	var nextPageToken string = ""
	var referenceImagesElement *visionpb.ReferenceImage = &visionpb.ReferenceImage{}
	var referenceImages = []*visionpb.ReferenceImage{referenceImagesElement}
	var expectedResponse = &visionpb.ListReferenceImagesResponse{
		PageSize:        pageSize,
		NextPageToken:   nextPageToken,
		ReferenceImages: referenceImages,
	}

	mockProductSearch.err = nil
	mockProductSearch.reqs = nil

	mockProductSearch.resps = append(mockProductSearch.resps[:0], expectedResponse)

	var formattedParent string = fmt.Sprintf("projects/%s/locations/%s/products/%s", "[PROJECT]", "[LOCATION]", "[PRODUCT]")
	var request = &visionpb.ListReferenceImagesRequest{
		Parent: formattedParent,
	}

	c, err := NewProductSearchClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := c.ListReferenceImages(context.Background(), request).Next()

	if err != nil {
		t.Fatal(err)
	}

	if want, got := request, mockProductSearch.reqs[0]; !proto.Equal(want, got) {
		t.Errorf("wrong request %q, want %q", got, want)
	}

	want := (interface{})(expectedResponse.ReferenceImages[0])
	got := (interface{})(resp)
	var ok bool

	switch want := (want).(type) {
	case proto.Message:
		ok = proto.Equal(want, got.(proto.Message))
	default:
		ok = want == got
	}
	if !ok {
		t.Errorf("wrong response %q, want %q)", got, want)
	}
}

func TestProductSearchListReferenceImagesError(t *testing.T) {
	errCode := codes.PermissionDenied
	mockProductSearch.err = gstatus.Error(errCode, "test error")

	var formattedParent string = fmt.Sprintf("projects/%s/locations/%s/products/%s", "[PROJECT]", "[LOCATION]", "[PRODUCT]")
	var request = &visionpb.ListReferenceImagesRequest{
		Parent: formattedParent,
	}

	c, err := NewProductSearchClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := c.ListReferenceImages(context.Background(), request).Next()

	if st, ok := gstatus.FromError(err); !ok {
		t.Errorf("got error %v, expected grpc error", err)
	} else if c := st.Code(); c != errCode {
		t.Errorf("got error code %q, want %q", c, errCode)
	}
	_ = resp
}
func TestProductSearchGetReferenceImage(t *testing.T) {
	var name2 string = "name2-1052831874"
	var uri string = "uri116076"
	var expectedResponse = &visionpb.ReferenceImage{
		Name: name2,
		Uri:  uri,
	}

	mockProductSearch.err = nil
	mockProductSearch.reqs = nil

	mockProductSearch.resps = append(mockProductSearch.resps[:0], expectedResponse)

	var formattedName string = fmt.Sprintf("projects/%s/locations/%s/products/%s/referenceImages/%s", "[PROJECT]", "[LOCATION]", "[PRODUCT]", "[IMAGE]")
	var request = &visionpb.GetReferenceImageRequest{
		Name: formattedName,
	}

	c, err := NewProductSearchClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := c.GetReferenceImage(context.Background(), request)

	if err != nil {
		t.Fatal(err)
	}

	if want, got := request, mockProductSearch.reqs[0]; !proto.Equal(want, got) {
		t.Errorf("wrong request %q, want %q", got, want)
	}

	if want, got := expectedResponse, resp; !proto.Equal(want, got) {
		t.Errorf("wrong response %q, want %q)", got, want)
	}
}

func TestProductSearchGetReferenceImageError(t *testing.T) {
	errCode := codes.PermissionDenied
	mockProductSearch.err = gstatus.Error(errCode, "test error")

	var formattedName string = fmt.Sprintf("projects/%s/locations/%s/products/%s/referenceImages/%s", "[PROJECT]", "[LOCATION]", "[PRODUCT]", "[IMAGE]")
	var request = &visionpb.GetReferenceImageRequest{
		Name: formattedName,
	}

	c, err := NewProductSearchClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := c.GetReferenceImage(context.Background(), request)

	if st, ok := gstatus.FromError(err); !ok {
		t.Errorf("got error %v, expected grpc error", err)
	} else if c := st.Code(); c != errCode {
		t.Errorf("got error code %q, want %q", c, errCode)
	}
	_ = resp
}
func TestProductSearchDeleteReferenceImage(t *testing.T) {
	var expectedResponse *emptypb.Empty = &emptypb.Empty{}

	mockProductSearch.err = nil
	mockProductSearch.reqs = nil

	mockProductSearch.resps = append(mockProductSearch.resps[:0], expectedResponse)

	var formattedName string = fmt.Sprintf("projects/%s/locations/%s/products/%s/referenceImages/%s", "[PROJECT]", "[LOCATION]", "[PRODUCT]", "[IMAGE]")
	var request = &visionpb.DeleteReferenceImageRequest{
		Name: formattedName,
	}

	c, err := NewProductSearchClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	err = c.DeleteReferenceImage(context.Background(), request)

	if err != nil {
		t.Fatal(err)
	}

	if want, got := request, mockProductSearch.reqs[0]; !proto.Equal(want, got) {
		t.Errorf("wrong request %q, want %q", got, want)
	}

}

func TestProductSearchDeleteReferenceImageError(t *testing.T) {
	errCode := codes.PermissionDenied
	mockProductSearch.err = gstatus.Error(errCode, "test error")

	var formattedName string = fmt.Sprintf("projects/%s/locations/%s/products/%s/referenceImages/%s", "[PROJECT]", "[LOCATION]", "[PRODUCT]", "[IMAGE]")
	var request = &visionpb.DeleteReferenceImageRequest{
		Name: formattedName,
	}

	c, err := NewProductSearchClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	err = c.DeleteReferenceImage(context.Background(), request)

	if st, ok := gstatus.FromError(err); !ok {
		t.Errorf("got error %v, expected grpc error", err)
	} else if c := st.Code(); c != errCode {
		t.Errorf("got error code %q, want %q", c, errCode)
	}
}
func TestProductSearchCreateReferenceImage(t *testing.T) {
	var name string = "name3373707"
	var uri string = "uri116076"
	var expectedResponse = &visionpb.ReferenceImage{
		Name: name,
		Uri:  uri,
	}

	mockProductSearch.err = nil
	mockProductSearch.reqs = nil

	mockProductSearch.resps = append(mockProductSearch.resps[:0], expectedResponse)

	var formattedParent string = fmt.Sprintf("projects/%s/locations/%s/products/%s", "[PROJECT]", "[LOCATION]", "[PRODUCT]")
	var referenceImage *visionpb.ReferenceImage = &visionpb.ReferenceImage{}
	var request = &visionpb.CreateReferenceImageRequest{
		Parent:         formattedParent,
		ReferenceImage: referenceImage,
	}

	c, err := NewProductSearchClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := c.CreateReferenceImage(context.Background(), request)

	if err != nil {
		t.Fatal(err)
	}

	if want, got := request, mockProductSearch.reqs[0]; !proto.Equal(want, got) {
		t.Errorf("wrong request %q, want %q", got, want)
	}

	if want, got := expectedResponse, resp; !proto.Equal(want, got) {
		t.Errorf("wrong response %q, want %q)", got, want)
	}
}

func TestProductSearchCreateReferenceImageError(t *testing.T) {
	errCode := codes.PermissionDenied
	mockProductSearch.err = gstatus.Error(errCode, "test error")

	var formattedParent string = fmt.Sprintf("projects/%s/locations/%s/products/%s", "[PROJECT]", "[LOCATION]", "[PRODUCT]")
	var referenceImage *visionpb.ReferenceImage = &visionpb.ReferenceImage{}
	var request = &visionpb.CreateReferenceImageRequest{
		Parent:         formattedParent,
		ReferenceImage: referenceImage,
	}

	c, err := NewProductSearchClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := c.CreateReferenceImage(context.Background(), request)

	if st, ok := gstatus.FromError(err); !ok {
		t.Errorf("got error %v, expected grpc error", err)
	} else if c := st.Code(); c != errCode {
		t.Errorf("got error code %q, want %q", c, errCode)
	}
	_ = resp
}
func TestProductSearchCreateProductSet(t *testing.T) {
	var name string = "name3373707"
	var displayName string = "displayName1615086568"
	var expectedResponse = &visionpb.ProductSet{
		Name:        name,
		DisplayName: displayName,
	}

	mockProductSearch.err = nil
	mockProductSearch.reqs = nil

	mockProductSearch.resps = append(mockProductSearch.resps[:0], expectedResponse)

	var formattedParent string = fmt.Sprintf("projects/%s/locations/%s", "[PROJECT]", "[LOCATION]")
	var productSet *visionpb.ProductSet = &visionpb.ProductSet{}
	var request = &visionpb.CreateProductSetRequest{
		Parent:     formattedParent,
		ProductSet: productSet,
	}

	c, err := NewProductSearchClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := c.CreateProductSet(context.Background(), request)

	if err != nil {
		t.Fatal(err)
	}

	if want, got := request, mockProductSearch.reqs[0]; !proto.Equal(want, got) {
		t.Errorf("wrong request %q, want %q", got, want)
	}

	if want, got := expectedResponse, resp; !proto.Equal(want, got) {
		t.Errorf("wrong response %q, want %q)", got, want)
	}
}

func TestProductSearchCreateProductSetError(t *testing.T) {
	errCode := codes.PermissionDenied
	mockProductSearch.err = gstatus.Error(errCode, "test error")

	var formattedParent string = fmt.Sprintf("projects/%s/locations/%s", "[PROJECT]", "[LOCATION]")
	var productSet *visionpb.ProductSet = &visionpb.ProductSet{}
	var request = &visionpb.CreateProductSetRequest{
		Parent:     formattedParent,
		ProductSet: productSet,
	}

	c, err := NewProductSearchClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := c.CreateProductSet(context.Background(), request)

	if st, ok := gstatus.FromError(err); !ok {
		t.Errorf("got error %v, expected grpc error", err)
	} else if c := st.Code(); c != errCode {
		t.Errorf("got error code %q, want %q", c, errCode)
	}
	_ = resp
}
func TestProductSearchListProductSets(t *testing.T) {
	var nextPageToken string = ""
	var productSetsElement *visionpb.ProductSet = &visionpb.ProductSet{}
	var productSets = []*visionpb.ProductSet{productSetsElement}
	var expectedResponse = &visionpb.ListProductSetsResponse{
		NextPageToken: nextPageToken,
		ProductSets:   productSets,
	}

	mockProductSearch.err = nil
	mockProductSearch.reqs = nil

	mockProductSearch.resps = append(mockProductSearch.resps[:0], expectedResponse)

	var formattedParent string = fmt.Sprintf("projects/%s/locations/%s", "[PROJECT]", "[LOCATION]")
	var request = &visionpb.ListProductSetsRequest{
		Parent: formattedParent,
	}

	c, err := NewProductSearchClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := c.ListProductSets(context.Background(), request).Next()

	if err != nil {
		t.Fatal(err)
	}

	if want, got := request, mockProductSearch.reqs[0]; !proto.Equal(want, got) {
		t.Errorf("wrong request %q, want %q", got, want)
	}

	want := (interface{})(expectedResponse.ProductSets[0])
	got := (interface{})(resp)
	var ok bool

	switch want := (want).(type) {
	case proto.Message:
		ok = proto.Equal(want, got.(proto.Message))
	default:
		ok = want == got
	}
	if !ok {
		t.Errorf("wrong response %q, want %q)", got, want)
	}
}

func TestProductSearchListProductSetsError(t *testing.T) {
	errCode := codes.PermissionDenied
	mockProductSearch.err = gstatus.Error(errCode, "test error")

	var formattedParent string = fmt.Sprintf("projects/%s/locations/%s", "[PROJECT]", "[LOCATION]")
	var request = &visionpb.ListProductSetsRequest{
		Parent: formattedParent,
	}

	c, err := NewProductSearchClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := c.ListProductSets(context.Background(), request).Next()

	if st, ok := gstatus.FromError(err); !ok {
		t.Errorf("got error %v, expected grpc error", err)
	} else if c := st.Code(); c != errCode {
		t.Errorf("got error code %q, want %q", c, errCode)
	}
	_ = resp
}
func TestProductSearchGetProductSet(t *testing.T) {
	var name2 string = "name2-1052831874"
	var displayName string = "displayName1615086568"
	var expectedResponse = &visionpb.ProductSet{
		Name:        name2,
		DisplayName: displayName,
	}

	mockProductSearch.err = nil
	mockProductSearch.reqs = nil

	mockProductSearch.resps = append(mockProductSearch.resps[:0], expectedResponse)

	var formattedName string = fmt.Sprintf("projects/%s/locations/%s/productSets/%s", "[PROJECT]", "[LOCATION]", "[PRODUCT_SET]")
	var request = &visionpb.GetProductSetRequest{
		Name: formattedName,
	}

	c, err := NewProductSearchClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := c.GetProductSet(context.Background(), request)

	if err != nil {
		t.Fatal(err)
	}

	if want, got := request, mockProductSearch.reqs[0]; !proto.Equal(want, got) {
		t.Errorf("wrong request %q, want %q", got, want)
	}

	if want, got := expectedResponse, resp; !proto.Equal(want, got) {
		t.Errorf("wrong response %q, want %q)", got, want)
	}
}

func TestProductSearchGetProductSetError(t *testing.T) {
	errCode := codes.PermissionDenied
	mockProductSearch.err = gstatus.Error(errCode, "test error")

	var formattedName string = fmt.Sprintf("projects/%s/locations/%s/productSets/%s", "[PROJECT]", "[LOCATION]", "[PRODUCT_SET]")
	var request = &visionpb.GetProductSetRequest{
		Name: formattedName,
	}

	c, err := NewProductSearchClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := c.GetProductSet(context.Background(), request)

	if st, ok := gstatus.FromError(err); !ok {
		t.Errorf("got error %v, expected grpc error", err)
	} else if c := st.Code(); c != errCode {
		t.Errorf("got error code %q, want %q", c, errCode)
	}
	_ = resp
}
func TestProductSearchUpdateProductSet(t *testing.T) {
	var name string = "name3373707"
	var displayName string = "displayName1615086568"
	var expectedResponse = &visionpb.ProductSet{
		Name:        name,
		DisplayName: displayName,
	}

	mockProductSearch.err = nil
	mockProductSearch.reqs = nil

	mockProductSearch.resps = append(mockProductSearch.resps[:0], expectedResponse)

	var productSet *visionpb.ProductSet = &visionpb.ProductSet{}
	var request = &visionpb.UpdateProductSetRequest{
		ProductSet: productSet,
	}

	c, err := NewProductSearchClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := c.UpdateProductSet(context.Background(), request)

	if err != nil {
		t.Fatal(err)
	}

	if want, got := request, mockProductSearch.reqs[0]; !proto.Equal(want, got) {
		t.Errorf("wrong request %q, want %q", got, want)
	}

	if want, got := expectedResponse, resp; !proto.Equal(want, got) {
		t.Errorf("wrong response %q, want %q)", got, want)
	}
}

func TestProductSearchUpdateProductSetError(t *testing.T) {
	errCode := codes.PermissionDenied
	mockProductSearch.err = gstatus.Error(errCode, "test error")

	var productSet *visionpb.ProductSet = &visionpb.ProductSet{}
	var request = &visionpb.UpdateProductSetRequest{
		ProductSet: productSet,
	}

	c, err := NewProductSearchClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := c.UpdateProductSet(context.Background(), request)

	if st, ok := gstatus.FromError(err); !ok {
		t.Errorf("got error %v, expected grpc error", err)
	} else if c := st.Code(); c != errCode {
		t.Errorf("got error code %q, want %q", c, errCode)
	}
	_ = resp
}
func TestProductSearchDeleteProductSet(t *testing.T) {
	var expectedResponse *emptypb.Empty = &emptypb.Empty{}

	mockProductSearch.err = nil
	mockProductSearch.reqs = nil

	mockProductSearch.resps = append(mockProductSearch.resps[:0], expectedResponse)

	var formattedName string = fmt.Sprintf("projects/%s/locations/%s/productSets/%s", "[PROJECT]", "[LOCATION]", "[PRODUCT_SET]")
	var request = &visionpb.DeleteProductSetRequest{
		Name: formattedName,
	}

	c, err := NewProductSearchClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	err = c.DeleteProductSet(context.Background(), request)

	if err != nil {
		t.Fatal(err)
	}

	if want, got := request, mockProductSearch.reqs[0]; !proto.Equal(want, got) {
		t.Errorf("wrong request %q, want %q", got, want)
	}

}

func TestProductSearchDeleteProductSetError(t *testing.T) {
	errCode := codes.PermissionDenied
	mockProductSearch.err = gstatus.Error(errCode, "test error")

	var formattedName string = fmt.Sprintf("projects/%s/locations/%s/productSets/%s", "[PROJECT]", "[LOCATION]", "[PRODUCT_SET]")
	var request = &visionpb.DeleteProductSetRequest{
		Name: formattedName,
	}

	c, err := NewProductSearchClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	err = c.DeleteProductSet(context.Background(), request)

	if st, ok := gstatus.FromError(err); !ok {
		t.Errorf("got error %v, expected grpc error", err)
	} else if c := st.Code(); c != errCode {
		t.Errorf("got error code %q, want %q", c, errCode)
	}
}
func TestProductSearchAddProductToProductSet(t *testing.T) {
	var expectedResponse *emptypb.Empty = &emptypb.Empty{}

	mockProductSearch.err = nil
	mockProductSearch.reqs = nil

	mockProductSearch.resps = append(mockProductSearch.resps[:0], expectedResponse)

	var formattedName string = fmt.Sprintf("projects/%s/locations/%s/productSets/%s", "[PROJECT]", "[LOCATION]", "[PRODUCT_SET]")
	var product string = "product-309474065"
	var request = &visionpb.AddProductToProductSetRequest{
		Name:    formattedName,
		Product: product,
	}

	c, err := NewProductSearchClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	err = c.AddProductToProductSet(context.Background(), request)

	if err != nil {
		t.Fatal(err)
	}

	if want, got := request, mockProductSearch.reqs[0]; !proto.Equal(want, got) {
		t.Errorf("wrong request %q, want %q", got, want)
	}

}

func TestProductSearchAddProductToProductSetError(t *testing.T) {
	errCode := codes.PermissionDenied
	mockProductSearch.err = gstatus.Error(errCode, "test error")

	var formattedName string = fmt.Sprintf("projects/%s/locations/%s/productSets/%s", "[PROJECT]", "[LOCATION]", "[PRODUCT_SET]")
	var product string = "product-309474065"
	var request = &visionpb.AddProductToProductSetRequest{
		Name:    formattedName,
		Product: product,
	}

	c, err := NewProductSearchClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	err = c.AddProductToProductSet(context.Background(), request)

	if st, ok := gstatus.FromError(err); !ok {
		t.Errorf("got error %v, expected grpc error", err)
	} else if c := st.Code(); c != errCode {
		t.Errorf("got error code %q, want %q", c, errCode)
	}
}
func TestProductSearchRemoveProductFromProductSet(t *testing.T) {
	var expectedResponse *emptypb.Empty = &emptypb.Empty{}

	mockProductSearch.err = nil
	mockProductSearch.reqs = nil

	mockProductSearch.resps = append(mockProductSearch.resps[:0], expectedResponse)

	var formattedName string = fmt.Sprintf("projects/%s/locations/%s/productSets/%s", "[PROJECT]", "[LOCATION]", "[PRODUCT_SET]")
	var product string = "product-309474065"
	var request = &visionpb.RemoveProductFromProductSetRequest{
		Name:    formattedName,
		Product: product,
	}

	c, err := NewProductSearchClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	err = c.RemoveProductFromProductSet(context.Background(), request)

	if err != nil {
		t.Fatal(err)
	}

	if want, got := request, mockProductSearch.reqs[0]; !proto.Equal(want, got) {
		t.Errorf("wrong request %q, want %q", got, want)
	}

}

func TestProductSearchRemoveProductFromProductSetError(t *testing.T) {
	errCode := codes.PermissionDenied
	mockProductSearch.err = gstatus.Error(errCode, "test error")

	var formattedName string = fmt.Sprintf("projects/%s/locations/%s/productSets/%s", "[PROJECT]", "[LOCATION]", "[PRODUCT_SET]")
	var product string = "product-309474065"
	var request = &visionpb.RemoveProductFromProductSetRequest{
		Name:    formattedName,
		Product: product,
	}

	c, err := NewProductSearchClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	err = c.RemoveProductFromProductSet(context.Background(), request)

	if st, ok := gstatus.FromError(err); !ok {
		t.Errorf("got error %v, expected grpc error", err)
	} else if c := st.Code(); c != errCode {
		t.Errorf("got error code %q, want %q", c, errCode)
	}
}
func TestProductSearchListProductsInProductSet(t *testing.T) {
	var nextPageToken string = ""
	var productsElement *visionpb.Product = &visionpb.Product{}
	var products = []*visionpb.Product{productsElement}
	var expectedResponse = &visionpb.ListProductsInProductSetResponse{
		NextPageToken: nextPageToken,
		Products:      products,
	}

	mockProductSearch.err = nil
	mockProductSearch.reqs = nil

	mockProductSearch.resps = append(mockProductSearch.resps[:0], expectedResponse)

	var formattedName string = fmt.Sprintf("projects/%s/locations/%s/productSets/%s", "[PROJECT]", "[LOCATION]", "[PRODUCT_SET]")
	var request = &visionpb.ListProductsInProductSetRequest{
		Name: formattedName,
	}

	c, err := NewProductSearchClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := c.ListProductsInProductSet(context.Background(), request).Next()

	if err != nil {
		t.Fatal(err)
	}

	if want, got := request, mockProductSearch.reqs[0]; !proto.Equal(want, got) {
		t.Errorf("wrong request %q, want %q", got, want)
	}

	want := (interface{})(expectedResponse.Products[0])
	got := (interface{})(resp)
	var ok bool

	switch want := (want).(type) {
	case proto.Message:
		ok = proto.Equal(want, got.(proto.Message))
	default:
		ok = want == got
	}
	if !ok {
		t.Errorf("wrong response %q, want %q)", got, want)
	}
}

func TestProductSearchListProductsInProductSetError(t *testing.T) {
	errCode := codes.PermissionDenied
	mockProductSearch.err = gstatus.Error(errCode, "test error")

	var formattedName string = fmt.Sprintf("projects/%s/locations/%s/productSets/%s", "[PROJECT]", "[LOCATION]", "[PRODUCT_SET]")
	var request = &visionpb.ListProductsInProductSetRequest{
		Name: formattedName,
	}

	c, err := NewProductSearchClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := c.ListProductsInProductSet(context.Background(), request).Next()

	if st, ok := gstatus.FromError(err); !ok {
		t.Errorf("got error %v, expected grpc error", err)
	} else if c := st.Code(); c != errCode {
		t.Errorf("got error code %q, want %q", c, errCode)
	}
	_ = resp
}
func TestProductSearchImportProductSets(t *testing.T) {
	var expectedResponse *visionpb.ImportProductSetsResponse = &visionpb.ImportProductSetsResponse{}

	mockProductSearch.err = nil
	mockProductSearch.reqs = nil

	any, err := ptypes.MarshalAny(expectedResponse)
	if err != nil {
		t.Fatal(err)
	}
	mockProductSearch.resps = append(mockProductSearch.resps[:0], &longrunningpb.Operation{
		Name:   "longrunning-test",
		Done:   true,
		Result: &longrunningpb.Operation_Response{Response: any},
	})

	var formattedParent string = fmt.Sprintf("projects/%s/locations/%s", "[PROJECT]", "[LOCATION]")
	var inputConfig *visionpb.ImportProductSetsInputConfig = &visionpb.ImportProductSetsInputConfig{}
	var request = &visionpb.ImportProductSetsRequest{
		Parent:      formattedParent,
		InputConfig: inputConfig,
	}

	c, err := NewProductSearchClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	respLRO, err := c.ImportProductSets(context.Background(), request)
	if err != nil {
		t.Fatal(err)
	}
	resp, err := respLRO.Wait(context.Background())

	if err != nil {
		t.Fatal(err)
	}

	if want, got := request, mockProductSearch.reqs[0]; !proto.Equal(want, got) {
		t.Errorf("wrong request %q, want %q", got, want)
	}

	if want, got := expectedResponse, resp; !proto.Equal(want, got) {
		t.Errorf("wrong response %q, want %q)", got, want)
	}
}

func TestProductSearchImportProductSetsError(t *testing.T) {
	errCode := codes.PermissionDenied
	mockProductSearch.err = nil
	mockProductSearch.resps = append(mockProductSearch.resps[:0], &longrunningpb.Operation{
		Name: "longrunning-test",
		Done: true,
		Result: &longrunningpb.Operation_Error{
			Error: &status.Status{
				Code:    int32(errCode),
				Message: "test error",
			},
		},
	})

	var formattedParent string = fmt.Sprintf("projects/%s/locations/%s", "[PROJECT]", "[LOCATION]")
	var inputConfig *visionpb.ImportProductSetsInputConfig = &visionpb.ImportProductSetsInputConfig{}
	var request = &visionpb.ImportProductSetsRequest{
		Parent:      formattedParent,
		InputConfig: inputConfig,
	}

	c, err := NewProductSearchClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	respLRO, err := c.ImportProductSets(context.Background(), request)
	if err != nil {
		t.Fatal(err)
	}
	resp, err := respLRO.Wait(context.Background())

	if st, ok := gstatus.FromError(err); !ok {
		t.Errorf("got error %v, expected grpc error", err)
	} else if c := st.Code(); c != errCode {
		t.Errorf("got error code %q, want %q", c, errCode)
	}
	_ = resp
}
func TestImageAnnotatorBatchAnnotateImages(t *testing.T) {
	var expectedResponse *visionpb.BatchAnnotateImagesResponse = &visionpb.BatchAnnotateImagesResponse{}

	mockImageAnnotator.err = nil
	mockImageAnnotator.reqs = nil

	mockImageAnnotator.resps = append(mockImageAnnotator.resps[:0], expectedResponse)

	var requests []*visionpb.AnnotateImageRequest = nil
	var request = &visionpb.BatchAnnotateImagesRequest{
		Requests: requests,
	}

	c, err := NewImageAnnotatorClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := c.BatchAnnotateImages(context.Background(), request)

	if err != nil {
		t.Fatal(err)
	}

	if want, got := request, mockImageAnnotator.reqs[0]; !proto.Equal(want, got) {
		t.Errorf("wrong request %q, want %q", got, want)
	}

	if want, got := expectedResponse, resp; !proto.Equal(want, got) {
		t.Errorf("wrong response %q, want %q)", got, want)
	}
}

func TestImageAnnotatorBatchAnnotateImagesError(t *testing.T) {
	errCode := codes.PermissionDenied
	mockImageAnnotator.err = gstatus.Error(errCode, "test error")

	var requests []*visionpb.AnnotateImageRequest = nil
	var request = &visionpb.BatchAnnotateImagesRequest{
		Requests: requests,
	}

	c, err := NewImageAnnotatorClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := c.BatchAnnotateImages(context.Background(), request)

	if st, ok := gstatus.FromError(err); !ok {
		t.Errorf("got error %v, expected grpc error", err)
	} else if c := st.Code(); c != errCode {
		t.Errorf("got error code %q, want %q", c, errCode)
	}
	_ = resp
}
func TestImageAnnotatorAsyncBatchAnnotateFiles(t *testing.T) {
	var expectedResponse *visionpb.AsyncBatchAnnotateFilesResponse = &visionpb.AsyncBatchAnnotateFilesResponse{}

	mockImageAnnotator.err = nil
	mockImageAnnotator.reqs = nil

	any, err := ptypes.MarshalAny(expectedResponse)
	if err != nil {
		t.Fatal(err)
	}
	mockImageAnnotator.resps = append(mockImageAnnotator.resps[:0], &longrunningpb.Operation{
		Name:   "longrunning-test",
		Done:   true,
		Result: &longrunningpb.Operation_Response{Response: any},
	})

	var requests []*visionpb.AsyncAnnotateFileRequest = nil
	var request = &visionpb.AsyncBatchAnnotateFilesRequest{
		Requests: requests,
	}

	c, err := NewImageAnnotatorClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	respLRO, err := c.AsyncBatchAnnotateFiles(context.Background(), request)
	if err != nil {
		t.Fatal(err)
	}
	resp, err := respLRO.Wait(context.Background())

	if err != nil {
		t.Fatal(err)
	}

	if want, got := request, mockImageAnnotator.reqs[0]; !proto.Equal(want, got) {
		t.Errorf("wrong request %q, want %q", got, want)
	}

	if want, got := expectedResponse, resp; !proto.Equal(want, got) {
		t.Errorf("wrong response %q, want %q)", got, want)
	}
}

func TestImageAnnotatorAsyncBatchAnnotateFilesError(t *testing.T) {
	errCode := codes.PermissionDenied
	mockImageAnnotator.err = nil
	mockImageAnnotator.resps = append(mockImageAnnotator.resps[:0], &longrunningpb.Operation{
		Name: "longrunning-test",
		Done: true,
		Result: &longrunningpb.Operation_Error{
			Error: &status.Status{
				Code:    int32(errCode),
				Message: "test error",
			},
		},
	})

	var requests []*visionpb.AsyncAnnotateFileRequest = nil
	var request = &visionpb.AsyncBatchAnnotateFilesRequest{
		Requests: requests,
	}

	c, err := NewImageAnnotatorClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	respLRO, err := c.AsyncBatchAnnotateFiles(context.Background(), request)
	if err != nil {
		t.Fatal(err)
	}
	resp, err := respLRO.Wait(context.Background())

	if st, ok := gstatus.FromError(err); !ok {
		t.Errorf("got error %v, expected grpc error", err)
	} else if c := st.Code(); c != errCode {
		t.Errorf("got error code %q, want %q", c, errCode)
	}
	_ = resp
}

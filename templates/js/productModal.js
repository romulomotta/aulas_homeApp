window.addEventListener('load', open())


function open() {
    console.log('teste pagina');
    const btnNewProduct = document.querySelector("#btn-new-product");
btnNewProduct.addEventListener('click', newProductModal());
}

function newProductModal(){

}

$(function(){
    $('#exampleModal').modal('show');
});
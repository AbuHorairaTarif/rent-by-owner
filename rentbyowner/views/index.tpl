{{template "partials/header.tpl" .}}
{{template "partials/nav.tpl" .}}

<div class= "container-fluid">
  <div class="row">
    <div class="col">
      <div aria-label="breadcrumb">
      <ol class="breadcrumb">
      Vacation Rental in &nbsp;
      <li class="breadcrumb-item"><a href="#">Bangladesh</a></li>
      <li class="breadcrumb-item active" aria-current="page">Chittagong</li>
      </ol>
      </div>
   </div>
  </div>  
  <div class="row">
  <h2 class="title">Chittagong Division Vacation Rentals & Rent By Owner Homes</h2>
  </div>
  <div class="row">
    <form>
    <button type="button" class="btn btn-outline-primary">Dates</button>
    <button type="button" class="btn btn-outline-primary">Prices</button>
    <button type="button" class="btn btn-outline-primary">Guests</button>
    <button type="button" class="btn btn-outline-primary">More Filters</button>
    <select class="form-select" aria-label="Default select example">
    <option selected>Select Option</option>
    <option value="1">Lowest Price</option>
    <option value="2">Highest Price</option>
    <option value="3">Most Popular</option>
  </select>
    </form>
  </div>
</div>



{{template "partials/footer.tpl" .}}
<div class="container-fluid mb-3 mt-5 ">
 
          <div class="row row-cols-1 row-cols-md-2 row-cols-lg-3 g-4 bg-card mx-2">
           {{ range.Stay.Data }}
              <div class="col">
                <div class="card h-90 shadow">
                  <div class="img-card">
                      <button type="button" class="btn btn-light btn-new">New</button>
                      <img src="https://cf.bstatic.com{{ .BasicPropertyData.Photos.Main.HighResJpegURL.RelativeURL}}" class="card-img-top"
                      alt="Hollywood Sign on The Hill" loading="lazy" />
                      <button type="button" class="btn btn-light btn-money">From {{.PriceDisplayInfoIrene.DisplayPrice.AmountPerStay.Amount}}</button>
                      <button type="button" class="btn btn-light opacity-75 btn-love"><i class="fa-regular fa-heart"></i></button>
                      <button type="button" class="btn btn-light opacity-75 btn-map"><i class="fa-solid fa-location-dot"></i></button>
                  </div>
                 
                  <div class="card-body">
                   
                    <p class="card-text">
        
                       <small> <i class="fa fa-star" aria-hidden="true"></i>{{.BasicPropertyData.StarRating.Value}}
                      ({{.BasicPropertyData.Reviews.ReviewsCount}} Reviews)</small>
                      <span class="col text-right text-end">
                            <small class="type-category">Apartment</small>
                          </span>
                     <a href="/property/details/?id_detail={{.IDDetail}}"><h6 class="card-title">{{.DisplayName.Title}}</h6></a>
                      <div class="col-auto amenities">
                            <span class="mt-2"><small><i class="fa-solid fa-circle"></i></small> View  <small><i class="fa-solid fa-circle"></i></small> Child Friendly</span>
                          </div>
                     
                      <small><a href="#" class="location-text">{{.BasicPropertyData.Location.City}} > {{.BasicPropertyData.Location.Address}}</a></small>
                     
                      <div class="text-end">
                        <span class="logo-img"><img src="https://static.rentbyowner.com/fa1eeed78042f188cf8bc13b2e6e4fcf898e2e18/static/images/booking.svg" alt="logo"></span>
                        <button class=" btn btn-view-availability">View Availability</button>
                      </div>
                      <div class="mt-3">
                        <div class="row">
                         
                         
                      </div>

                      </div>
                    </p>
                  </div>
                </div>
               
              </div>
              {{end}}
</div>
              
</div>
document.addEventListener('DOMContentLoaded', function() {
    // Dropdown menu functionality
    var dropdowns = document.querySelectorAll('.dropdown-toggle');

    dropdowns.forEach(function(dropdown) {
        dropdown.addEventListener('click', function(event) {
            event.preventDefault();
            var menu = this.nextElementSibling;
            if (menu.style.display === 'block') {
                menu.style.display = 'none';
            } else {
                menu.style.display = 'block';
            }
        });
    });

    document.addEventListener('click', function(event) {
        dropdowns.forEach(function(dropdown) {
            var menu = dropdown.nextElementSibling;
            if (!dropdown.contains(event.target) && menu.style.display === 'block') {
                menu.style.display = 'none';
            }
        });
    });

    // Slider functionality
    var sliderContainer = document.querySelector('.slider');
    var slides = [];
    var currentSlide = 0;

    // Preload images
    function preloadImages(imagePaths, callback) {
        var loadedImages = 0;
        imagePaths.forEach(function(path, index) {
            var img = new Image();
            img.src = path;
            img.onload = function() {
                loadedImages++;
                slides[index] = img;
                if (loadedImages === imagePaths.length) {
                    callback();
                }
            };
        });
    }

    // Change slide
    function changeSlide() {
        sliderContainer.innerHTML = '';
        sliderContainer.appendChild(slides[currentSlide]);
        currentSlide = (currentSlide + 1) % slides.length;
    }

    // Get slide paths and initialize slider
    fetch('/images/slides')
        .then(response => response.json())
        .then(data => {
            var imagePaths = data.map(filename => `/images/slides/${filename}`);
            preloadImages(imagePaths, function() {
                changeSlide();
                setInterval(changeSlide, 3000); // Change slide every 3 seconds
            });
        })
        .catch(error => console.error('Error fetching slide images:', error));
});

(() => {
  const modal = document.querySelector("[data-gallery-modal]");
  if (!modal) {
    return;
  }

  const frame = modal.querySelector("[data-gallery-frame]");
  const count = modal.querySelector("[data-gallery-count]");
  const caption = modal.querySelector("[data-gallery-caption]");
  const prevButton = modal.querySelector("[data-gallery-prev]");
  const nextButton = modal.querySelector("[data-gallery-next]");
  const closeButton = modal.querySelector(".gallery-modal-close");

  let slides = [];
  let activeIndex = 0;
  let lastTrigger = null;

  function collectSlides(group) {
    return Array.from(document.querySelectorAll("[data-gallery-trigger]"))
      .filter((element) => element.dataset.galleryGroup === group)
      .map((element) => ({
        src: element.dataset.gallerySrc || "",
        alt: element.dataset.galleryAlt || "",
        caption: element.dataset.galleryCaption || "",
      }));
  }

  function renderSlide() {
    if (!slides.length) {
      return;
    }

    const slide = slides[activeIndex];
    frame.replaceChildren();

    if (slide.src) {
      const image = document.createElement("img");
      image.className = "gallery-modal-image";
      image.src = slide.src;
      image.alt = slide.alt || slide.caption || "Gallery image";
      image.loading = "eager";
      frame.append(image);
    } else {
      const placeholder = document.createElement("div");
      placeholder.className = "gallery-modal-placeholder";
      placeholder.textContent = slide.caption || slide.alt || "Photo";
      frame.append(placeholder);
    }

    count.textContent = `${activeIndex + 1} / ${slides.length}`;
    caption.textContent = slide.caption || slide.alt || "";

    const hasMultipleSlides = slides.length > 1;
    prevButton.disabled = !hasMultipleSlides;
    nextButton.disabled = !hasMultipleSlides;
  }

  function openGallery(trigger) {
    const group = trigger.dataset.galleryGroup;
    slides = collectSlides(group);
    activeIndex = Number(trigger.dataset.galleryIndex || "0");
    lastTrigger = trigger;

    if (!slides.length) {
      return;
    }

    renderSlide();
    modal.hidden = false;
    document.body.classList.add("gallery-open");
    closeButton.focus();
  }

  function closeGallery() {
    modal.hidden = true;
    document.body.classList.remove("gallery-open");
    frame.replaceChildren();
    slides = [];
    count.textContent = "";
    caption.textContent = "";

    if (lastTrigger) {
      lastTrigger.focus();
      lastTrigger = null;
    }
  }

  function moveGallery(step) {
    if (slides.length < 2) {
      return;
    }

    activeIndex = (activeIndex + step + slides.length) % slides.length;
    renderSlide();
  }

  document.addEventListener("click", (event) => {
    const trigger = event.target.closest("[data-gallery-trigger]");
    if (trigger) {
      event.preventDefault();
      openGallery(trigger);
      return;
    }

    if (!modal.hidden && event.target.closest("[data-gallery-close]")) {
      event.preventDefault();
      closeGallery();
      return;
    }

    if (!modal.hidden && event.target.closest("[data-gallery-prev]")) {
      event.preventDefault();
      moveGallery(-1);
      return;
    }

    if (!modal.hidden && event.target.closest("[data-gallery-next]")) {
      event.preventDefault();
      moveGallery(1);
    }
  });

  document.addEventListener("keydown", (event) => {
    if (modal.hidden) {
      return;
    }

    if (event.key === "Escape") {
      closeGallery();
      return;
    }

    if (event.key === "ArrowLeft") {
      moveGallery(-1);
      return;
    }

    if (event.key === "ArrowRight") {
      moveGallery(1);
    }
  });
})();

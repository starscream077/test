const reveals = document.querySelectorAll('.reveal');

const observer = new IntersectionObserver(entries => {
  entries.forEach(entry => {
    if (entry.isIntersecting) {
      entry.target.classList.add('visible');
    }
  });
});

reveals.forEach(r => observer.observe(r));

const toggle = document.getElementById('themeToggle');
let neon = false;

toggle.onclick = () => {
  neon = !neon;
  document.documentElement.style.setProperty(
    '--accent',
    neon ? '#22d3ee' : '#38bdf8'
  );
};

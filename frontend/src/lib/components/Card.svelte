<script>
    export let src;
    export let alt;
</script>

<div class="group relative m-1">
    <div class="absolute -inset-0.5 bg-gradient-to-r from-purple-500 to-pink-500 rounded-lg blur opacity-0 group-hover:opacity-100 transition duration-300"></div>
    <div class="w-24 h-32 relative transform transition-all duration-300 hover:scale-110 hover:-translate-y-1 cursor-pointer">
        <img {src} alt={alt} class="w-full h-full object-contain rounded-lg shadow-xl"/>
        <div class="absolute inset-0 rounded-lg bg-black opacity-0 group-hover:opacity-10 transition-all duration-300"></div>
    </div>
</div>
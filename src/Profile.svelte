<script>
    import { Link } from "svelte-routing";
</script>

<main>
    <h1>Local events <span>Near You</span></h1>
    <div class="scan__box">
        <h2>Join Event</h2>
        <div class="scan__cam"></div>
        <button class="solid">SCAN QR</button>
    </div>
    <div class="btn__box">
        <Link to="/events-register">
            <button class="solid">REGISTER EVENTS</button>
        </Link>
    </div>
</main>

<style>
    .scan__cam {
        height: 50vh;
        background-color: var(--color-light);
        margin: 0.5rem 0 2rem 0;
    }
</style>

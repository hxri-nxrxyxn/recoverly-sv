<script>
    import { Link } from "svelte-routing";
</script>

<main>
    <h1>Local events <br /><span>Near You</span></h1>

    <div class="card">
        <div class="card__image">
            <img src="#" alt="" />
        </div>
        <div class="card__text">
            <h1>Anti Drug Campaign</h1>
            <p>National Service Scheme</p>
            <p>Location</p>
        </div>
        <Link to="/event">
            <button>REGISTER</button>
        </Link>
    </div>
    <div class="card">
        <div class="card__image">
            <img src="#" alt="" />
        </div>
        <div class="card__text">
            <h1>ExSmoker Summit</h1>
            <p>Instagram Support</p>
            <p>Location</p>
        </div>
        <Link to="/event">
            <button>REGISTER</button>
        </Link>
    </div>
    <div class="card">
        <div class="card__image">
            <img src="#" alt="" />
        </div>
        <div class="card__text">
            <h1>Anti Drug Campaign</h1>
            <p>National Service Scheme</p>
            <p>Location</p>
        </div>
        <Link to="/event">
            <button>REGISTER</button>
        </Link>
    </div>
</main>

<style>
    .card {
        margin-bottom: 2rem;
    }
</style>

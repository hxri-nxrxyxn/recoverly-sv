<script>
    import { Link } from "svelte-routing";
</script>

<main>
    <img src="#" alt="" />
    <div class="details">
        <h1>Anti-drug campaign</h1>
        <h2>National Service Scheme</h2>
        <p>
            Lorem, ipsum dolor sit amet consectetur adipisicing elit. Sit harum
            corrupti officia expedita vel! Tempore obcaecati ad unde sint
            blanditiis recusandae ullam, neque consectetur nostrum. Iusto fugiat
            tempora sint sapiente.
        </p>
        <div class="details__line">
            <h3>Niagra Falls</h3>
            <p><b>10AM</b> 11 - 09 - 2024</p>
        </div>
    </div>
    <div class="btn__box">
        <Link to="/terms">
            <button class="solid">REGISTER</button>
        </Link>
    </div>
</main>

<style>
    .details__line {
        display: flex;
        justify-content: space-between;
    }
</style>

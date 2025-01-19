<script>
    import Star from "./assets/noun-star-923969.svg";
    import ProfilePic from "./assets/circle-user.svg";
    import { Link } from "svelte-routing";

    const viewDropDown = () => {
        document.querySelector(".nav__dropdown").style.display = "block";
    };
    const hideDropDown = () => {
        document.querySelector(".nav__dropdown").style.display = "none";
    };
</script>

<main>
    <div class="nav">
        <div class="nav__icon">
            <img src={Star} alt="" class="icon" onclick={viewDropDown} />
        </div>
        <div class="nav__dropdown">
            <p onclick={hideDropDown}>Home</p>
            <Link to="/chatroom">
                <p>Chatroom</p>
            </Link>
            <Link to="/events">
                <p>Local Events</p>
            </Link>
            <Link>
                <p>Templates & Guides</p>
            </Link>
            <Link>
                <p>Service</p>
            </Link>
            <Link>
                <p>AI Task Generation</p>
            </Link>
        </div>
        <div class="nav__placeholder">hey</div>
        <div class="nav__stats">
            <h5>Hari</h5>
            <img src={ProfilePic} alt="" class="icon" />
        </div>
    </div>
</main>

<style>
    .nav {
        display: flex;
        background-color: white;
        z-index: 99;
        margin-bottom: 0.5rem;
        position: relative;
    }
    .nav__placeholder {
        flex: 1;
        color: transparent;
    }
    .nav__icon img {
        scale: 1.2;
    }
    .nav__stats {
        display: flex;
        align-items: flex-start;
        gap: 0.5rem;
    }
    .nav__dropdown {
        position: absolute;
        left: 0;
        top: 0;
        background-color: white;
        width: 100%;
        height: 100vh;
        display: none;
    }
    .nav__dropdown p {
        font-size: 1.2rem;
        margin-bottom: 1rem;
        font-weight: 300;
    }
</style>

<script>
    import { Link } from "svelte-routing";
</script>

<main>
    <div class="pills">
        <div class="pill">
            <p>LEVEL</p>
        </div>
        <div class="pill">
            <p>EVENTS</p>
        </div>
        <div class="pill">
            <p>CHATROOM</p>
        </div>
    </div>
    <h1>Good morning, <span>Hari</span></h1>

    <div class="card">
        <h1>Daily Task</h1>
        <div class="tasks">
            <div class="task">
                <div class="task__text">
                    <h5>No. of minutes meditated</h5>
                    <p>0/15</p>
                </div>
                <div class="task__adjust">
                    <button>-</button>
                    <h5>0</h5>
                    <button>+</button>
                </div>
            </div>
            <div class="task">
                <div class="task__text">
                    <h5>No. of cigarretes smoked</h5>
                    <p>0/7</p>
                </div>
                <div class="task__adjust">
                    <button>-</button>
                    <h5>0</h5>
                    <button>+</button>
                </div>
            </div>
            <button>submit</button>
        </div>
    </div>

    <div class="achievements">
        <div class="achievement">
            <div class="achievement__logo">
                <i>logo</i>
            </div>
            <div class="achievement__text">
                <p>Fresh Start</p>
            </div>
        </div>
        <div class="achievement">
            <div class="achievement__logo">
                <i>logo</i>
            </div>
            <div class="achievement__text">
                <p>Fresh Start</p>
            </div>
        </div>
    </div>

    <div class="emergency">
        <Link to="/emergency">SOS</Link>
    </div>
</main>

<style>
    .task {
        display: flex;
        justify-content: space-between;
        margin-top: 1rem;
    }
    .task__adjust {
        display: flex;
        align-items: center;
        justify-content: space-between;
    }
    .task__adjust > * {
        margin-left: 1rem;
    }
    .task__adjust button {
        position: initial;
        font-size: 1rem;
        padding: 0.5rem;
        border: 0;
        color: var(--color-text);
        background-color: var(--color-light);
    }
    .task__adjust h5 {
        text-align: center;
    }
    .task__text p {
        font-size: 0.75rem;
    }
    .pills {
        display: flex;
        margin-bottom: 1rem;
    }
    .pill {
        background-color: var(--color-light);
        padding: 0.25rem 1rem;
        margin-right: 0.5rem;
        border-radius: 2rem;
    }
    .pill p {
        font-size: 0.5rem;
        font-weight: 700;
        color: var(--color-primary);
    }
    .achievement {
        display: flex;
    }
</style>

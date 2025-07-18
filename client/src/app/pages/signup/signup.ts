import { Component, ElementRef, ViewChild } from '@angular/core';
import { Router, RouterLink } from '@angular/router';
import { FormsModule } from '@angular/forms';
import { CommonModule } from '@angular/common';
import { AuthService } from '../../services/auth/auth';
import { ToastrService } from 'ngx-toastr';
import { InputField } from '../../components/input-field/input-field';
import { animate, style, transition, trigger } from '@angular/animations';

@Component({
  selector: 'app-signup',
  standalone: true,
  imports: [RouterLink, FormsModule, CommonModule, InputField],
  templateUrl: './signup.html',
  styleUrl: './signup.css',
  animations: [
    trigger('fadeIn', [
      transition(':enter', [ // when element enters the DOM
        style({ opacity: 0, transform: 'translateY(10px)' }),
        animate('300ms ease-out', style({ opacity: 1, transform: 'translateY(0)' }))
      ])
    ])
  ]
})
export class Signup {
  email = '';
  password = '';
  confirmPassword = '';
  message = '';
  error = '';
  showElement = false;
  @ViewChild('submitButton') submitButtonRef!: ElementRef<HTMLButtonElement>;

  constructor(private auth: AuthService, private router: Router, private toastr: ToastrService) {}

  ngAfterViewInit() {
    setTimeout(() => {
      this.showElement = true;
    }, 100);
  }

  disableButton() {
    if (this.submitButtonRef.nativeElement) {
      this.submitButtonRef.nativeElement.disabled = true; // Disable button to prevent multiple clicks
    }
  }

  enableButton() {
    if (this.submitButtonRef.nativeElement) {
      this.submitButtonRef.nativeElement.disabled = false; // Enable button
    }
  }

  onSignupSubmit() {
    this.disableButton();

    this.auth.signup({
      email: this.email,
      password: this.password,
      confirm_password: this.confirmPassword,
    }).subscribe({
      next: () => {
        this.toastr.success('Account created successfully!', 'Success');
        setTimeout(() => this.router.navigate(['/login']), 1500);
      },
      error: (err) => {
        if (this.password !== this.confirmPassword) {
          this.toastr.error('Passwords do not match.', 'Error');
        }
        else {
          this.toastr.error(err?.error?.message || 'Signup failed. Please try again.', 'Error');
        }
        this.enableButton(); // Re-enable button on error
      }
    });

    
  }
}
